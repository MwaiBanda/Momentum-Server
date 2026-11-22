package application

import (
	httpHandler "Momentum/infrastructure/http"
	healthHTTP "Momentum/infrastructure/http/health"
	httpMiddleware "Momentum/infrastructure/http/middleware"

	"context"
	"time"

	"github.com/gofiber/contrib/v3/uptime"
	"github.com/gofiber/fiber/v3"
	fiberredis "github.com/gofiber/storage/redis/v3"
)

func (app *Application) configureHTTP(handlers *httpHandler.Handler) {
	app.server = fiber.New(fiber.Config{ErrorHandler: httpHandler.ErrorHandler})
	httpMiddleware.ConfigureAppMiddleware(app.server)
	app.server.Use("/metrics", httpMiddleware.BearerTokenAuthentication(app.config.Security.APIKey))
	app.server.Use(httpMiddleware.Prometheus(app.config.Telemetry.ServiceName))

	healthHandlers := healthHTTP.New(
		func(ctx context.Context) error {
			return app.database.PingContext(ctx)
		},
		func(ctx context.Context) error {
			return app.redisClient.Ping(ctx).Err()
		},
	)
	app.server.Get("/healthz", healthHandlers.Liveness)
	app.server.Get("/readyz", healthHandlers.Readiness)

	app.uptimeStore = fiberredis.NewFromConnection(app.redisClient)
	if app.config.Uptime.RequireBearer {
		app.server.Use("/uptime", httpMiddleware.BearerTokenAuthentication(app.config.Security.APIKey))
	}
	app.server.Use(uptime.New(uptime.Config{
		App:              app.server,
		ServiceID:        app.config.Uptime.ServiceID,
		ServiceName:      app.config.Uptime.ServiceName,
		Store:            app.uptimeStore,
		StorageKeyPrefix: app.config.Uptime.KeyPrefix,
		RetentionDays:    app.config.Uptime.RetentionDays,
		SampleInterval:   time.Duration(app.config.Uptime.SampleIntervalS) * time.Second,
	}))

	app.registerRoutes(handlers)
}

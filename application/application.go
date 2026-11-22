package application

import (
	appconfig "Momentum/application/config"
	database "Momentum/infrastructure/persistence/db"

	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gofiber/fiber/v3"
	fiberredis "github.com/gofiber/storage/redis/v3"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

const shutdownTimeout = 10 * time.Second

type Application struct {
	config            *appconfig.Config
	server            *fiber.App
	shutdownTelemetry func(context.Context) error
	database          *sqlx.DB
	redisClient       *redis.Client
	uptimeStore       *fiberredis.Storage
	closeOnce         sync.Once
	closeErr          error
}

func New(ctx context.Context) (_ *Application, err error) {
	app := &Application{}
	defer func() {
		if err != nil {
			if cleanupErr := app.Close(); cleanupErr != nil {
				err = errors.Join(err, fmt.Errorf("cleanup failed initialization: %w", cleanupErr))
			}
		}
	}()

	if err = app.initialize(ctx); err != nil {
		return nil, err
	}
	return app, nil
}

func (app *Application) Run(ctx context.Context) error {
	if err := ctx.Err(); err != nil {
		return nil
	}
	if err := app.server.Listen(
		app.config.Host+":"+app.config.Port,
		fiber.ListenConfig{
			GracefulContext: ctx,
			ShutdownTimeout: shutdownTimeout,
		},
	); err != nil {
		return fmt.Errorf("listen: %w", err)
	}
	return nil
}

func (app *Application) Close() error {
	app.closeOnce.Do(func() {
		var closeErrors []error

		if app.uptimeStore != nil {
			if err := app.uptimeStore.Close(); err != nil {
				closeErrors = append(closeErrors, fmt.Errorf("close Fiber Uptime store: %w", err))
			}
		}
		if app.redisClient != nil {
			if err := app.redisClient.Close(); err != nil {
				closeErrors = append(closeErrors, fmt.Errorf("close Redis client: %w", err))
			}
		}
		if app.database != nil {
			if err := database.Close(app.database); err != nil {
				closeErrors = append(closeErrors, fmt.Errorf("close SQL database: %w", err))
			}
		}
		if app.shutdownTelemetry != nil {
			ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
			defer cancel()
			if err := app.shutdownTelemetry(ctx); err != nil {
				log.Printf("shutdown OpenTelemetry: %v", err)
			}
		}

		app.closeErr = errors.Join(closeErrors...)
	})
	return app.closeErr
}

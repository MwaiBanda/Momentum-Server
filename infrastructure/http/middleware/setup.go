package middleware

import (
	"Momentum/infrastructure/http/request"

	fiberotel "github.com/gofiber/contrib/v3/otel"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func ConfigureAppMiddleware(app *fiber.App) {
	app.Use(recover.New())
	app.Use(request.RequestID())
	app.Use(fiberotel.New(
		fiberotel.WithoutMetrics(true),
		fiberotel.WithClientIP(false),
		fiberotel.WithRedactedQueryParams("token", "api_key", "access_token"),
	))
	configureCors(app)
	configureLogger(app)
	configureHelmet(app)
	configureFiles(app)
}

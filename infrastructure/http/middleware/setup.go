package middleware

import (
	"Momentum/infrastructure/http/request"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
)

func ConfigureAppMiddleware(app *fiber.App) {
	app.Use(recover.New())
	app.Use(request.RequestID())
	configureCors(app)
	configureLogger(app)
	configureHelmet(app)
	configureFiles(app)
}

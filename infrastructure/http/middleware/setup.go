package middleware

import "github.com/gofiber/fiber/v3"

func ConfigureAppMiddleware(app *fiber.App) {
	configureCors(app)
	configureLogger(app)
	configureHelmet(app)
	configureFiles(app)
}

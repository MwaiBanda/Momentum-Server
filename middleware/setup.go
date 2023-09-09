package middleware

import "github.com/gofiber/fiber/v2"

func ConfigureAppMiddleWare(app *fiber.App) {
	configureCors(app)
	configureLogger(app)
	configureHelmet(app)
	configureFiles(app)
}

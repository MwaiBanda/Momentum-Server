package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func configureHelmet(app *fiber.App) {
	app.Use(helmet.New())
}

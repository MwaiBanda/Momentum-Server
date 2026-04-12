package middleware

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/helmet"
)

func configureHelmet(app *fiber.App) {
	app.Use(helmet.New())
}

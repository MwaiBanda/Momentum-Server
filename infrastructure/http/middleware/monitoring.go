package middleware

import (
	"github.com/gofiber/contrib/v3/monitor"
	"github.com/gofiber/fiber/v3"
)

func Monitor() fiber.Handler {
	return monitor.New()
}

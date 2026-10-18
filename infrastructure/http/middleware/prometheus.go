package middleware

import (
	fiberprometheus "github.com/gofiber/contrib/v3/prometheus"
	"github.com/gofiber/fiber/v3"
)

func Prometheus(serviceName string) fiber.Handler {
	return fiberprometheus.New(fiberprometheus.Config{
		ServiceName: serviceName,
		MetricsPath: "/metrics",
	})
}

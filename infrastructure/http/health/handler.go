package health

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
)

const defaultReadinessTimeout = 2 * time.Second

type Check func(context.Context) error

type Handler struct {
	checks           map[string]Check
	readinessTimeout time.Duration
}

type response struct {
	Status string            `json:"status"`
	Checks map[string]string `json:"checks,omitempty"`
}

func New(databaseCheck, redisCheck Check) *Handler {
	return &Handler{
		checks: map[string]Check{
			"database": databaseCheck,
			"redis":    redisCheck,
		},
		readinessTimeout: defaultReadinessTimeout,
	}
}

func (handler *Handler) Liveness(ctx fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(response{Status: "ok"})
}

func (handler *Handler) Readiness(ctx fiber.Ctx) error {
	checkContext, cancel := context.WithTimeout(ctx.Context(), handler.readinessTimeout)
	defer cancel()

	checks := make(map[string]string, len(handler.checks))
	ready := true
	for name, check := range handler.checks {
		if check == nil {
			checks[name] = "failed"
			ready = false
			log.Printf("readiness check %s is not configured", name)
			continue
		}

		if err := check(checkContext); err != nil {
			checks[name] = "failed"
			ready = false
			log.Printf("readiness check %s failed: %v", name, err)
			continue
		}
		checks[name] = "ok"
	}

	status := fiber.StatusOK
	statusText := "ready"
	if !ready {
		status = fiber.StatusServiceUnavailable
		statusText = "not_ready"
	}

	return ctx.Status(status).JSON(response{
		Status: statusText,
		Checks: checks,
	})
}

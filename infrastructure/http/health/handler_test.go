package health

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestLivenessIsProcessOnly(t *testing.T) {
	app := fiber.New()
	handler := New(
		func(context.Context) error { return errors.New("database unavailable") },
		func(context.Context) error { return errors.New("redis unavailable") },
	)
	app.Get("/healthz", handler.Liveness)

	response, err := app.Test(httptest.NewRequest(http.MethodGet, "/healthz", nil))
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	if response.StatusCode != fiber.StatusOK {
		t.Fatalf("status = %d, want %d", response.StatusCode, fiber.StatusOK)
	}
}

func TestReadinessReportsDependencyFailure(t *testing.T) {
	app := fiber.New()
	handler := New(
		func(context.Context) error { return nil },
		func(context.Context) error { return errors.New("redis unavailable") },
	)
	app.Get("/readyz", handler.Readiness)

	response, err := app.Test(httptest.NewRequest(http.MethodGet, "/readyz", nil))
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	if response.StatusCode != fiber.StatusServiceUnavailable {
		t.Fatalf("status = %d, want %d", response.StatusCode, fiber.StatusServiceUnavailable)
	}
}

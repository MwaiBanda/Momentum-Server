package request

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestRequestIDPreservesSafeIncomingID(t *testing.T) {
	app := fiber.New()
	app.Use(RequestID())
	app.Get("/", func(ctx fiber.Ctx) error {
		if got := FromContext(ctx); got != "request-123" {
			t.Fatalf("request ID in context = %q, want request-123", got)
		}
		return ctx.SendStatus(http.StatusNoContent)
	})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	request.Header.Set(RequestIDHeader, "request-123")
	response, err := app.Test(request)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer response.Body.Close()

	if got := response.Header.Get(RequestIDHeader); got != "request-123" {
		t.Fatalf("response request ID = %q, want request-123", got)
	}
}

func TestRequestIDGeneratesWhenIncomingIDIsInvalid(t *testing.T) {
	app := fiber.New()
	app.Use(RequestID())
	app.Get("/", func(ctx fiber.Ctx) error {
		if FromContext(ctx) == "" {
			t.Fatal("request ID was not stored in context")
		}
		return ctx.SendStatus(http.StatusNoContent)
	})

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	request.Header.Set(RequestIDHeader, "unsafe request id")
	response, err := app.Test(request)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer response.Body.Close()

	if got := response.Header.Get(RequestIDHeader); got == "" || got == "unsafe request id" {
		t.Fatalf("generated request ID = %q, want a safe generated value", got)
	}
}

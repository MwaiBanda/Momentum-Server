package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestErrorHandlerReturnsUnifiedResponse(t *testing.T) {
	app := fiber.New(fiber.Config{ErrorHandler: ErrorHandler})
	app.Get("/", func(fiber.Ctx) error {
		return fiber.ErrNotFound
	})

	response, err := app.Test(httptest.NewRequest(http.MethodGet, "/", nil))
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer response.Body.Close()

	var body ErrorResponse
	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		t.Fatalf("decode response = %v", err)
	}

	if response.StatusCode != http.StatusNotFound {
		t.Fatalf("status = %d, want %d", response.StatusCode, http.StatusNotFound)
	}
	if body != (ErrorResponse{Code: http.StatusNotFound, Message: http.StatusText(http.StatusNotFound)}) {
		t.Fatalf("body = %#v, want unified not found response", body)
	}
	if !strings.HasPrefix(response.Header.Get("Content-Type"), fiber.MIMEApplicationJSON) {
		t.Fatalf("content type = %q, want %q", response.Header.Get("Content-Type"), fiber.MIMEApplicationJSON)
	}
}

func TestErrorHandlerHidesUnexpectedError(t *testing.T) {
	app := fiber.New(fiber.Config{ErrorHandler: ErrorHandler})
	app.Get("/", func(fiber.Ctx) error {
		return errors.New("database password")
	})

	response, err := app.Test(httptest.NewRequest(http.MethodGet, "/", nil))
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer response.Body.Close()

	var body ErrorResponse
	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		t.Fatalf("decode response = %v", err)
	}

	if response.StatusCode != http.StatusInternalServerError {
		t.Fatalf("status = %d, want %d", response.StatusCode, http.StatusInternalServerError)
	}
	if body.Message != http.StatusText(http.StatusInternalServerError) {
		t.Fatalf("message = %q, want generic internal server error", body.Message)
	}
}

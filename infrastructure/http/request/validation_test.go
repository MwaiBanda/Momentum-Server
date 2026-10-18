package request

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v3"
)

type testRequest struct {
	Name    string            `json:"name" binding:"required"`
	Enabled bool              `json:"enabled" binding:"required"`
	Count   int               `json:"count" binding:"required"`
	Items   []testRequestItem `json:"items" binding:"required"`
}

type testRequestItem struct {
	ID string `json:"id" binding:"required"`
}

func TestBindAndValidateAcceptsExplicitZeroValues(t *testing.T) {
	app := fiber.New()
	app.Post("/", func(ctx fiber.Ctx) error {
		input, err := BindAndValidate[testRequest](ctx)
		if err != nil {
			return err
		}
		if input.Enabled || input.Count != 0 {
			t.Fatalf("input = %#v, want explicit zero values", input)
		}
		return ctx.SendStatus(http.StatusNoContent)
	})

	request := httptest.NewRequest(
		http.MethodPost,
		"/",
		strings.NewReader(`{"name":"","enabled":false,"count":0,"items":[{"id":""}]}`),
	)
	request.Header.Set("Content-Type", "application/json")
	response, err := app.Test(request)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusNoContent {
		t.Fatalf("status = %d, want %d", response.StatusCode, http.StatusNoContent)
	}
}

func TestBindAndValidateRejectsMissingNestedRequiredField(t *testing.T) {
	app := fiber.New(fiber.Config{ErrorHandler: func(ctx fiber.Ctx, err error) error {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}})
	app.Post("/", func(ctx fiber.Ctx) error {
		_, err := BindAndValidate[testRequest](ctx)
		return err
	})

	request := httptest.NewRequest(
		http.MethodPost,
		"/",
		strings.NewReader(`{"name":"request","enabled":true,"count":1,"items":[{}]}`),
	)
	request.Header.Set("Content-Type", "application/json")
	response, err := app.Test(request)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", response.StatusCode, http.StatusBadRequest)
	}
}

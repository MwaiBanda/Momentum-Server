package middleware

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v3"
)

func TestPrometheusMetricsRequireBearerAuthentication(t *testing.T) {
	app := fiber.New()
	app.Use("/metrics", BearerTokenAuthentication("secret"))
	app.Use(Prometheus("test"))
	app.Get("/healthz", func(ctx fiber.Ctx) error {
		return ctx.SendStatus(http.StatusOK)
	})

	unauthorizedResponse, err := app.Test(httptest.NewRequest(http.MethodGet, "/metrics", nil))
	if err != nil {
		t.Fatalf("unauthorized app.Test() error = %v", err)
	}
	if unauthorizedResponse.StatusCode != http.StatusUnauthorized {
		t.Fatalf("unauthorized status = %d, want %d", unauthorizedResponse.StatusCode, http.StatusUnauthorized)
	}

	if _, err := app.Test(httptest.NewRequest(http.MethodGet, "/healthz", nil)); err != nil {
		t.Fatalf("health app.Test() error = %v", err)
	}

	authorizedRequest := httptest.NewRequest(http.MethodGet, "/metrics", nil)
	authorizedRequest.Header.Set("Authorization", "Bearer secret")
	authorizedResponse, err := app.Test(authorizedRequest)
	if err != nil {
		t.Fatalf("authorized app.Test() error = %v", err)
	}
	if authorizedResponse.StatusCode != http.StatusOK {
		t.Fatalf("authorized status = %d, want %d", authorizedResponse.StatusCode, http.StatusOK)
	}

	body, err := io.ReadAll(authorizedResponse.Body)
	if err != nil {
		t.Fatalf("read metrics body: %v", err)
	}
	if !strings.Contains(string(body), "http_requests_total") {
		t.Fatalf("metrics body does not contain http_requests_total: %s", body)
	}
}

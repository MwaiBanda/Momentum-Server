package notification

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	applicationnotification "Momentum/application/service/notification"

	"github.com/gofiber/fiber/v3"
)

type fakeUseCase struct {
	send func(context.Context, applicationnotification.Input) (applicationnotification.Response, error)
}

func (fake *fakeUseCase) Send(ctx context.Context, input applicationnotification.Input) (applicationnotification.Response, error) {
	if fake.send == nil {
		return applicationnotification.Response{}, nil
	}
	return fake.send(ctx, input)
}

func newTestApp(notifications applicationnotification.UseCase) *fiber.App {
	handler := GetInstance(notifications)
	app := fiber.New()
	app.Post("/notifications", handler.Send)
	return app
}

func TestSendReturnsResponse(t *testing.T) {
	var received applicationnotification.Input
	expected := applicationnotification.Input{Title: "Title", Body: "Body", Topic: "all-users"}
	app := newTestApp(&fakeUseCase{
		send: func(_ context.Context, input applicationnotification.Input) (applicationnotification.Response, error) {
			received = input
			return applicationnotification.Response{Title: input.Title, Body: input.Body, Topic: input.Topic}, nil
		},
	})

	request := httptest.NewRequest(
		http.MethodPost,
		"/notifications",
		strings.NewReader(`{"title":"Title","body":"Body","topic":"all-users"}`),
	)
	request.Header.Set("Content-Type", "application/json")
	response, err := app.Test(request)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want %d", response.StatusCode, http.StatusOK)
	}
	if received != expected {
		t.Fatalf("input = %#v, want %#v", received, expected)
	}
}

func TestSendRejectsMalformedBody(t *testing.T) {
	called := false
	app := newTestApp(&fakeUseCase{
		send: func(context.Context, applicationnotification.Input) (applicationnotification.Response, error) {
			called = true
			return applicationnotification.Response{}, nil
		},
	})

	request := httptest.NewRequest(http.MethodPost, "/notifications", strings.NewReader(`{invalid`))
	request.Header.Set("Content-Type", "application/json")
	response, err := app.Test(request)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", response.StatusCode, http.StatusBadRequest)
	}
	if called {
		t.Fatal("Send use case was called for malformed input")
	}
}

func TestSendMapsSenderFailureToInternalServerError(t *testing.T) {
	app := newTestApp(&fakeUseCase{
		send: func(context.Context, applicationnotification.Input) (applicationnotification.Response, error) {
			return applicationnotification.Response{}, errors.New("firebase unavailable")
		},
	})

	request := httptest.NewRequest(http.MethodPost, "/notifications", strings.NewReader(`{"title":"Title","body":"Body","topic":"all-users"}`))
	request.Header.Set("Content-Type", "application/json")
	response, err := app.Test(request)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusInternalServerError {
		t.Fatalf("status = %d, want %d", response.StatusCode, http.StatusInternalServerError)
	}
}

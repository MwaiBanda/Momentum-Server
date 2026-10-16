package event

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	applicationevent "Momentum/application/service/event"

	"github.com/gofiber/fiber/v3"
)

type fakeUseCase struct {
	list func(context.Context) (applicationevent.Response, error)
}

func (fake *fakeUseCase) List(ctx context.Context) (applicationevent.Response, error) {
	if fake.list == nil {
		return applicationevent.Response{}, nil
	}
	return fake.list(ctx)
}

func TestListReturnsResponse(t *testing.T) {
	called := false
	app := fiber.New()
	app.Get("/events", GetInstance(&fakeUseCase{
		list: func(context.Context) (applicationevent.Response, error) {
			called = true
			return applicationevent.Response{
				Data: []applicationevent.Event{{ID: "event-1"}},
			}, nil
		},
	}).List)

	response, err := app.Test(httptest.NewRequest(http.MethodGet, "/events", nil))
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want %d", response.StatusCode, http.StatusOK)
	}
	if !called {
		t.Fatal("List use case was not called")
	}
}

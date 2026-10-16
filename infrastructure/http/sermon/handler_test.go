package sermon

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	applicationsermon "Momentum/application/service/sermon"

	"github.com/gofiber/fiber/v3"
)

type fakeUseCase struct {
	list func(context.Context, string) (applicationsermon.Response, error)
}

func (fake *fakeUseCase) List(ctx context.Context, page string) (applicationsermon.Response, error) {
	if fake.list == nil {
		return applicationsermon.Response{}, nil
	}
	return fake.list(ctx, page)
}

func TestListPassesPageAndReturnsResponse(t *testing.T) {
	var receivedPage string
	app := fiber.New()
	app.Get("/sermons", GetInstance(&fakeUseCase{
		list: func(_ context.Context, page string) (applicationsermon.Response, error) {
			receivedPage = page
			return applicationsermon.Response{CurrentPage: 2}, nil
		},
	}).List)

	response, err := app.Test(httptest.NewRequest(http.MethodGet, "/sermons?page=2", nil))
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want %d", response.StatusCode, http.StatusOK)
	}
	if receivedPage != "2" {
		t.Fatalf("page = %q, want %q", receivedPage, "2")
	}
}

func TestListMapsInvalidPageToBadRequest(t *testing.T) {
	app := fiber.New()
	app.Get("/sermons", GetInstance(&fakeUseCase{
		list: func(context.Context, string) (applicationsermon.Response, error) {
			return applicationsermon.Response{}, applicationsermon.ErrInvalidPage
		},
	}).List)

	response, err := app.Test(httptest.NewRequest(http.MethodGet, "/sermons?page=bad", nil))
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", response.StatusCode, http.StatusBadRequest)
	}
}

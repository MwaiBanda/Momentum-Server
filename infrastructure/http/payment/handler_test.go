package payment

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	applicationpayment "Momentum/application/service/payment"

	"github.com/gofiber/fiber/v3"
)

type fakeUseCase struct {
	create func(context.Context, applicationpayment.Input) (applicationpayment.Response, error)
}

func (fake *fakeUseCase) Create(ctx context.Context, input applicationpayment.Input) (applicationpayment.Response, error) {
	return fake.create(ctx, input)
}

func TestCreateRejectsMalformedBody(t *testing.T) {
	app := fiber.New()
	app.Post("/payments", GetInstance(&fakeUseCase{
		create: func(context.Context, applicationpayment.Input) (applicationpayment.Response, error) {
			t.Fatal("Create use case was called")
			return applicationpayment.Response{}, nil
		},
	}).Create)

	request := httptest.NewRequest(http.MethodPost, "/payments", strings.NewReader(`{invalid`))
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

func TestCreateMapsInvalidInput(t *testing.T) {
	app := fiber.New()
	app.Post("/payments", GetInstance(&fakeUseCase{
		create: func(context.Context, applicationpayment.Input) (applicationpayment.Response, error) {
			return applicationpayment.Response{}, applicationpayment.ErrInvalidInput
		},
	}).Create)

	response, err := app.Test(httptest.NewRequest(http.MethodPost, "/payments", strings.NewReader(`{}`)))
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", response.StatusCode, http.StatusBadRequest)
	}
}

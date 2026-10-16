package transaction

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	applicationtransaction "Momentum/application/service/transaction"

	"github.com/gofiber/fiber/v3"
)

type fakeUseCase struct {
	list       func(context.Context) ([]applicationtransaction.Response, error)
	listByUser func(context.Context, string) ([]applicationtransaction.Response, error)
	create     func(context.Context, applicationtransaction.Input) (applicationtransaction.Response, error)
	delete     func(context.Context, string) (applicationtransaction.Response, error)
}

func (fake *fakeUseCase) List(ctx context.Context) ([]applicationtransaction.Response, error) {
	return fake.list(ctx)
}

func (fake *fakeUseCase) ListByUser(ctx context.Context, userID string) ([]applicationtransaction.Response, error) {
	return fake.listByUser(ctx, userID)
}

func (fake *fakeUseCase) Create(ctx context.Context, input applicationtransaction.Input) (applicationtransaction.Response, error) {
	return fake.create(ctx, input)
}

func (fake *fakeUseCase) Delete(ctx context.Context, id string) (applicationtransaction.Response, error) {
	return fake.delete(ctx, id)
}

func TestCreateRejectsMalformedBody(t *testing.T) {
	app := fiber.New()
	app.Post("/transactions", GetInstance(&fakeUseCase{
		create: func(context.Context, applicationtransaction.Input) (applicationtransaction.Response, error) {
			t.Fatal("Create use case was called")
			return applicationtransaction.Response{}, nil
		},
	}).Create)

	request := httptest.NewRequest(http.MethodPost, "/transactions", strings.NewReader(`{invalid`))
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

func TestDeleteMapsErrors(t *testing.T) {
	app := fiber.New()
	app.Delete("/transactions/:transactionId", GetInstance(&fakeUseCase{
		delete: func(context.Context, string) (applicationtransaction.Response, error) {
			return applicationtransaction.Response{}, applicationtransaction.ErrNotFound
		},
	}).Delete)

	response, err := app.Test(httptest.NewRequest(http.MethodDelete, "/transactions/transaction-1", nil))
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusNotFound {
		t.Fatalf("status = %d, want %d", response.StatusCode, http.StatusNotFound)
	}
}

func TestListByUserPassesUserID(t *testing.T) {
	var receivedUserID string
	app := fiber.New()
	app.Get("/transactions/:userId", GetInstance(&fakeUseCase{
		listByUser: func(_ context.Context, userID string) ([]applicationtransaction.Response, error) {
			receivedUserID = userID
			return []applicationtransaction.Response{}, nil
		},
	}).ListByUser)

	response, err := app.Test(httptest.NewRequest(http.MethodGet, "/transactions/user-1", nil))
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK || receivedUserID != "user-1" {
		t.Fatalf("status = %d, user ID = %q", response.StatusCode, receivedUserID)
	}
}

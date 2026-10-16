package user

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	userService "Momentum/application/service/user"

	"github.com/gofiber/fiber/v3"
)

type fakeUseCase struct {
	getByID func(context.Context, string) (userService.Response, error)
	create  func(context.Context, userService.CreateInput) (userService.Response, error)
	update  func(context.Context, userService.UpdateInput) (userService.Response, error)
	delete  func(context.Context, string) (userService.Response, error)
}

func (fake *fakeUseCase) GetByID(ctx context.Context, id string) (userService.Response, error) {
	if fake.getByID == nil {
		return userService.Response{}, nil
	}
	return fake.getByID(ctx, id)
}

func (fake *fakeUseCase) Create(ctx context.Context, input userService.CreateInput) (userService.Response, error) {
	if fake.create == nil {
		return userService.Response{}, nil
	}
	return fake.create(ctx, input)
}

func (fake *fakeUseCase) Update(ctx context.Context, input userService.UpdateInput) (userService.Response, error) {
	if fake.update == nil {
		return userService.Response{}, nil
	}
	return fake.update(ctx, input)
}

func (fake *fakeUseCase) Delete(ctx context.Context, id string) (userService.Response, error) {
	if fake.delete == nil {
		return userService.Response{}, nil
	}
	return fake.delete(ctx, id)
}

func newTestApp(users userService.UseCase) *fiber.App {
	handler := GetInstance(users)
	app := fiber.New()
	app.Get("/users/:userId", handler.GetByID)
	app.Post("/users", handler.Create)
	app.Put("/users", handler.Update)
	app.Delete("/users/:userId", handler.Delete)
	return app
}

func performRequest(t *testing.T, app *fiber.App, method string, target string, body string) *http.Response {
	t.Helper()

	request := httptest.NewRequest(method, target, strings.NewReader(body))
	if body != "" {
		request.Header.Set("Content-Type", "application/json")
	}

	response, err := app.Test(request)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	t.Cleanup(func() {
		response.Body.Close()
	})

	return response
}

func TestGetByIDReturnsResponse(t *testing.T) {
	expected := userService.Response{
		ID:        "user-1",
		Email:     "user@example.com",
		Fullname:  "Test User",
		Phone:     "555-0100",
		CreatedOn: "2026-01-01T00:00:00Z",
	}
	var receivedID string
	users := &fakeUseCase{
		getByID: func(_ context.Context, id string) (userService.Response, error) {
			receivedID = id
			return expected, nil
		},
	}

	response := performRequest(t, newTestApp(users), http.MethodGet, "/users/user-1", "")
	if response.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want %d", response.StatusCode, http.StatusOK)
	}
	if receivedID != "user-1" {
		t.Fatalf("user ID = %q, want %q", receivedID, "user-1")
	}

	var actual userService.Response
	if err := json.NewDecoder(response.Body).Decode(&actual); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("response = %#v, want %#v", actual, expected)
	}
}

func TestGetByIDMapsErrors(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want int
	}{
		{name: "not found", err: userService.ErrNotFound, want: http.StatusNotFound},
		{name: "internal error", err: errors.New("database unavailable"), want: http.StatusInternalServerError},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			users := &fakeUseCase{
				getByID: func(context.Context, string) (userService.Response, error) {
					return userService.Response{}, test.err
				},
			}

			response := performRequest(t, newTestApp(users), http.MethodGet, "/users/user-1", "")
			if response.StatusCode != test.want {
				t.Fatalf("status = %d, want %d", response.StatusCode, test.want)
			}
		})
	}
}

func TestCreateReturnsResponse(t *testing.T) {
	input := userService.CreateInput{
		ID:       "user-1",
		Email:    "user@example.com",
		Fullname: "Test User",
		Phone:    "555-0100",
	}
	expected := userService.Response{ID: input.ID, Email: input.Email}
	var receivedInput userService.CreateInput
	users := &fakeUseCase{
		create: func(_ context.Context, actual userService.CreateInput) (userService.Response, error) {
			receivedInput = actual
			return expected, nil
		},
	}

	response := performRequest(t, newTestApp(users), http.MethodPost, "/users", `{"id":"user-1","email":"user@example.com","fullname":"Test User","phone":"555-0100"}`)
	if response.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want %d", response.StatusCode, http.StatusOK)
	}
	if !reflect.DeepEqual(receivedInput, input) {
		t.Fatalf("input = %#v, want %#v", receivedInput, input)
	}
}

func TestCreateRejectsMalformedBody(t *testing.T) {
	called := false
	users := &fakeUseCase{
		create: func(context.Context, userService.CreateInput) (userService.Response, error) {
			called = true
			return userService.Response{}, nil
		},
	}

	response := performRequest(t, newTestApp(users), http.MethodPost, "/users", `{invalid`)
	if response.StatusCode != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", response.StatusCode, http.StatusBadRequest)
	}
	if called {
		t.Fatal("Create use case was called for malformed input")
	}
}

func TestUpdateReturnsResponse(t *testing.T) {
	input := userService.UpdateInput{
		ID:       "user-1",
		Email:    "updated@example.com",
		Fullname: "Updated User",
		Phone:    "555-0101",
	}
	expected := userService.Response{ID: input.ID, Email: input.Email}
	var receivedInput userService.UpdateInput
	users := &fakeUseCase{
		update: func(_ context.Context, actual userService.UpdateInput) (userService.Response, error) {
			receivedInput = actual
			return expected, nil
		},
	}

	response := performRequest(t, newTestApp(users), http.MethodPut, "/users", `{"id":"user-1","email":"updated@example.com","fullname":"Updated User","phone":"555-0101"}`)
	if response.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want %d", response.StatusCode, http.StatusOK)
	}
	if !reflect.DeepEqual(receivedInput, input) {
		t.Fatalf("input = %#v, want %#v", receivedInput, input)
	}
}

func TestUpdateRejectsMalformedBody(t *testing.T) {
	called := false
	users := &fakeUseCase{
		update: func(context.Context, userService.UpdateInput) (userService.Response, error) {
			called = true
			return userService.Response{}, nil
		},
	}

	response := performRequest(t, newTestApp(users), http.MethodPut, "/users", `{invalid`)
	if response.StatusCode != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", response.StatusCode, http.StatusBadRequest)
	}
	if called {
		t.Fatal("Update use case was called for malformed input")
	}
}

func TestDeleteReturnsResponse(t *testing.T) {
	expected := userService.Response{ID: "user-1", Email: "deleted@example.com"}
	var receivedID string
	users := &fakeUseCase{
		delete: func(_ context.Context, id string) (userService.Response, error) {
			receivedID = id
			return expected, nil
		},
	}

	response := performRequest(t, newTestApp(users), http.MethodDelete, "/users/user-1", "")
	if response.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want %d", response.StatusCode, http.StatusOK)
	}
	if receivedID != "user-1" {
		t.Fatalf("user ID = %q, want %q", receivedID, "user-1")
	}
}

func TestDeleteMapsErrors(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want int
	}{
		{name: "not found", err: userService.ErrNotFound, want: http.StatusNotFound},
		{name: "internal error", err: errors.New("database unavailable"), want: http.StatusInternalServerError},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			users := &fakeUseCase{
				delete: func(context.Context, string) (userService.Response, error) {
					return userService.Response{}, test.err
				},
			}

			response := performRequest(t, newTestApp(users), http.MethodDelete, "/users/user-1", "")
			if response.StatusCode != test.want {
				t.Fatalf("status = %d, want %d", response.StatusCode, test.want)
			}
		})
	}
}

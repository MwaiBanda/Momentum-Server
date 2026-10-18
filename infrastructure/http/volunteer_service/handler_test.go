package volunteer_service

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	applicationservice "Momentum/application/service/volunteer_service"

	"github.com/gofiber/fiber/v3"
)

type fakeUseCase struct {
	list      func(context.Context) ([]applicationservice.ServiceResponse, error)
	getByType func(context.Context, string) (applicationservice.ServiceDetailResponse, error)
	create    func(context.Context, applicationservice.CreateInput) (applicationservice.CreateResponse, error)
	updateDay func(context.Context, applicationservice.UpdateDayInput) (applicationservice.UpdateDayResponse, error)
}

func (fake *fakeUseCase) ListServices(ctx context.Context) ([]applicationservice.ServiceResponse, error) {
	if fake.list == nil {
		return nil, nil
	}
	return fake.list(ctx)
}

func (fake *fakeUseCase) GetByType(ctx context.Context, serviceType string) (applicationservice.ServiceDetailResponse, error) {
	if fake.getByType == nil {
		return applicationservice.ServiceDetailResponse{}, nil
	}
	return fake.getByType(ctx, serviceType)
}

func (fake *fakeUseCase) Create(ctx context.Context, input applicationservice.CreateInput) (applicationservice.CreateResponse, error) {
	if fake.create == nil {
		return applicationservice.CreateResponse{}, nil
	}
	return fake.create(ctx, input)
}

func (fake *fakeUseCase) UpdateDay(ctx context.Context, input applicationservice.UpdateDayInput) (applicationservice.UpdateDayResponse, error) {
	if fake.updateDay == nil {
		return applicationservice.UpdateDayResponse{}, nil
	}
	return fake.updateDay(ctx, input)
}

func newTestApp(services applicationservice.UseCase) *fiber.App {
	handler := GetInstance(services)
	app := fiber.New()
	app.Get("/services", handler.ListServices)
	app.Get("/services/:type", handler.GetByType)
	app.Post("/services", handler.Create)
	app.Put("/services/day", handler.UpdateDay)
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

func TestListServicesReturnsResponse(t *testing.T) {
	expected := []applicationservice.ServiceResponse{{ID: "service-1", Name: "Food", Type: "food", Image: "food.png"}}
	services := &fakeUseCase{
		list: func(context.Context) ([]applicationservice.ServiceResponse, error) {
			return expected, nil
		},
	}

	response := performRequest(t, newTestApp(services), http.MethodGet, "/services", "")
	if response.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want %d", response.StatusCode, http.StatusOK)
	}

	var actual []applicationservice.ServiceResponse
	if err := json.NewDecoder(response.Body).Decode(&actual); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("response = %#v, want %#v", actual, expected)
	}
}

func TestGetByTypeMapsErrors(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want int
	}{
		{name: "invalid type", err: applicationservice.ErrInvalidInput, want: http.StatusBadRequest},
		{name: "not found", err: applicationservice.ErrServiceNotFound, want: http.StatusNotFound},
		{name: "internal error", err: errors.New("database unavailable"), want: http.StatusInternalServerError},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			services := &fakeUseCase{
				getByType: func(context.Context, string) (applicationservice.ServiceDetailResponse, error) {
					return applicationservice.ServiceDetailResponse{}, test.err
				},
			}

			response := performRequest(t, newTestApp(services), http.MethodGet, "/services/food", "")
			if response.StatusCode != test.want {
				t.Fatalf("status = %d, want %d", response.StatusCode, test.want)
			}
		})
	}
}

func TestCreateBindsRequest(t *testing.T) {
	input := applicationservice.CreateInput{
		ID:          "request-id",
		Organizer:   "Organizer",
		Description: "Description",
		Email:       "organizer@example.com",
		Phone:       "555-0100",
		Days:        []applicationservice.CreateDayInput{{Date: "2026-04-01"}},
		UserID:      "user-1",
		ServiceID:   "service-1",
	}
	var received applicationservice.CreateInput
	services := &fakeUseCase{
		create: func(_ context.Context, actual applicationservice.CreateInput) (applicationservice.CreateResponse, error) {
			received = actual
			return applicationservice.CreateResponse{
				ID:          actual.ID,
				Organizer:   actual.Organizer,
				Description: actual.Description,
				Email:       actual.Email,
				Phone:       actual.Phone,
				Days:        actual.Days,
				UserID:      actual.UserID,
				ServiceID:   actual.ServiceID,
			}, nil
		},
	}

	response := performRequest(
		t,
		newTestApp(services),
		http.MethodPost,
		"/services",
		`{"id":"request-id","organizer":"Organizer","description":"Description","email":"organizer@example.com","phone":"555-0100","days":[{"date":"2026-04-01"}],"user_id":"user-1","service_id":"service-1"}`,
	)
	if response.StatusCode != http.StatusOK {
		t.Fatalf("status = %d, want %d", response.StatusCode, http.StatusOK)
	}
	if !reflect.DeepEqual(received, input) {
		t.Fatalf("input = %#v, want %#v", received, input)
	}

	var actual applicationservice.CreateResponse
	if err := json.NewDecoder(response.Body).Decode(&actual); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	expected := applicationservice.CreateResponse{
		ID:          input.ID,
		Organizer:   input.Organizer,
		Description: input.Description,
		Email:       input.Email,
		Phone:       input.Phone,
		Days:        input.Days,
		UserID:      input.UserID,
		ServiceID:   input.ServiceID,
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("response = %#v, want %#v", actual, expected)
	}
}

func TestCreateRejectsMalformedBody(t *testing.T) {
	called := false
	services := &fakeUseCase{
		create: func(context.Context, applicationservice.CreateInput) (applicationservice.CreateResponse, error) {
			called = true
			return applicationservice.CreateResponse{}, nil
		},
	}

	response := performRequest(t, newTestApp(services), http.MethodPost, "/services", `{invalid`)
	if response.StatusCode != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", response.StatusCode, http.StatusBadRequest)
	}
	if called {
		t.Fatal("Create use case was called for malformed input")
	}
}

func TestUpdateDayMapsErrors(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want int
	}{
		{name: "invalid input", err: applicationservice.ErrInvalidInput, want: http.StatusBadRequest},
		{name: "not found", err: applicationservice.ErrDayNotFound, want: http.StatusNotFound},
		{name: "internal error", err: errors.New("database unavailable"), want: http.StatusInternalServerError},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			services := &fakeUseCase{
				updateDay: func(context.Context, applicationservice.UpdateDayInput) (applicationservice.UpdateDayResponse, error) {
					return applicationservice.UpdateDayResponse{}, test.err
				},
			}

			response := performRequest(t, newTestApp(services), http.MethodPut, "/services/day", `{"id":"day-1"}`)
			if response.StatusCode != test.want {
				t.Fatalf("status = %d, want %d", response.StatusCode, test.want)
			}
		})
	}
}

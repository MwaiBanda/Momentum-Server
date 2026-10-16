package meal

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	applicationmeal "Momentum/application/service/meal"

	"github.com/gofiber/fiber/v3"
)

type fakeUseCase struct {
	list                  func(context.Context) ([]applicationmeal.Response, error)
	create                func(context.Context, applicationmeal.Input) (applicationmeal.Input, error)
	update                func(context.Context, applicationmeal.Input) (applicationmeal.Input, error)
	delete                func(context.Context, string) (applicationmeal.Response, error)
	assignVolunteer       func(context.Context, applicationmeal.VolunteerAssignmentInput) (applicationmeal.VolunteerAssignmentInput, error)
	updateVolunteeredMeal func(context.Context, applicationmeal.UpdateVolunteeredMealInput) (applicationmeal.UpdateVolunteeredMealInput, error)
}

func (fake *fakeUseCase) List(ctx context.Context) ([]applicationmeal.Response, error) {
	return fake.list(ctx)
}

func (fake *fakeUseCase) Create(ctx context.Context, input applicationmeal.Input) (applicationmeal.Input, error) {
	return fake.create(ctx, input)
}

func (fake *fakeUseCase) Update(ctx context.Context, input applicationmeal.Input) (applicationmeal.Input, error) {
	return fake.update(ctx, input)
}

func (fake *fakeUseCase) Delete(ctx context.Context, id string) (applicationmeal.Response, error) {
	return fake.delete(ctx, id)
}

func (fake *fakeUseCase) AssignVolunteer(ctx context.Context, input applicationmeal.VolunteerAssignmentInput) (applicationmeal.VolunteerAssignmentInput, error) {
	return fake.assignVolunteer(ctx, input)
}

func (fake *fakeUseCase) UpdateVolunteeredMeal(ctx context.Context, input applicationmeal.UpdateVolunteeredMealInput) (applicationmeal.UpdateVolunteeredMealInput, error) {
	return fake.updateVolunteeredMeal(ctx, input)
}

func TestCreateRejectsMalformedBody(t *testing.T) {
	app := fiber.New()
	app.Post("/meals", GetInstance(&fakeUseCase{
		create: func(context.Context, applicationmeal.Input) (applicationmeal.Input, error) {
			t.Fatal("Create use case was called")
			return applicationmeal.Input{}, nil
		},
	}).Create)

	request := httptest.NewRequest(http.MethodPost, "/meals", strings.NewReader(`{invalid`))
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

func TestDeleteMapsNotFound(t *testing.T) {
	app := fiber.New()
	app.Delete("/meals/:mealId", GetInstance(&fakeUseCase{
		delete: func(context.Context, string) (applicationmeal.Response, error) {
			return applicationmeal.Response{}, applicationmeal.ErrNotFound
		},
	}).Delete)

	response, err := app.Test(httptest.NewRequest(http.MethodDelete, "/meals/meal-1", nil))
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusNotFound {
		t.Fatalf("status = %d, want %d", response.StatusCode, http.StatusNotFound)
	}
}

func TestAssignVolunteerPassesRequest(t *testing.T) {
	var received applicationmeal.VolunteerAssignmentInput
	app := fiber.New()
	app.Post("/meals/meal", GetInstance(&fakeUseCase{
		assignVolunteer: func(_ context.Context, input applicationmeal.VolunteerAssignmentInput) (applicationmeal.VolunteerAssignmentInput, error) {
			received = input
			return input, nil
		},
	}).AssignVolunteer)

	request := httptest.NewRequest(
		http.MethodPost,
		"/meals/meal",
		strings.NewReader(`{"meal_id":"meal-1","volunteered_meal":{"id":"volunteered-1","user":{"id":"user-1"}}}`),
	)
	request.Header.Set("Content-Type", "application/json")
	response, err := app.Test(request)
	if err != nil {
		t.Fatalf("app.Test() error = %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK || received.MealID != "meal-1" || received.VolunteeredMeal.User.ID != "user-1" {
		t.Fatalf("status = %d, received = %#v", response.StatusCode, received)
	}
}

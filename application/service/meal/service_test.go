package meal

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"testing"
	"time"

	"Momentum/application/cache"
	domainmeal "Momentum/domain/meal"
	domainuser "Momentum/domain/user"
)

type fakeRepository struct {
	meals                  []domainmeal.Meal
	createdParams          domainmeal.CreateParams
	updatedParams          domainmeal.UpdateParams
	assignParams           domainmeal.AssignVolunteerParams
	updatedVolunteerParams domainmeal.UpdateVolunteeredMealParams
	createErr              error
	updateErr              error
	deleteErr              error
	assignErr              error
	updateVolunteerErr     error
}

func (repository *fakeRepository) List(context.Context) ([]domainmeal.Meal, error) {
	return repository.meals, nil
}

func (repository *fakeRepository) Create(_ context.Context, params domainmeal.CreateParams) (domainmeal.Meal, error) {
	repository.createdParams = params
	return domainmeal.Meal{}, repository.createErr
}

func (repository *fakeRepository) Update(_ context.Context, params domainmeal.UpdateParams) (domainmeal.Meal, error) {
	repository.updatedParams = params
	return domainmeal.Meal{}, repository.updateErr
}

func (repository *fakeRepository) Delete(context.Context, string) (domainmeal.Meal, error) {
	return domainmeal.Meal{}, repository.deleteErr
}

func (repository *fakeRepository) AssignVolunteer(_ context.Context, params domainmeal.AssignVolunteerParams) (domainmeal.Meal, error) {
	repository.assignParams = params
	return domainmeal.Meal{}, repository.assignErr
}

func (repository *fakeRepository) UpdateVolunteeredMeal(_ context.Context, params domainmeal.UpdateVolunteeredMealParams) (domainmeal.VolunteeredMeal, error) {
	repository.updatedVolunteerParams = params
	return domainmeal.VolunteeredMeal{}, repository.updateVolunteerErr
}

type fakeCache struct {
	value     []byte
	found     bool
	setKey    string
	setTTL    time.Duration
	deleteKey string
}

func (cacheStore *fakeCache) Get(context.Context, string) ([]byte, bool, error) {
	return cacheStore.value, cacheStore.found, nil
}

func (cacheStore *fakeCache) Set(_ context.Context, key string, value []byte, ttl time.Duration) error {
	cacheStore.setKey = key
	cacheStore.value = value
	cacheStore.setTTL = ttl
	cacheStore.found = true
	return nil
}

func (cacheStore *fakeCache) Delete(_ context.Context, key string) error {
	cacheStore.deleteKey = key
	cacheStore.found = false
	return nil
}

func TestCreatePreservesRequestResponseAndGeneratesPersistenceIDs(t *testing.T) {
	repository := &fakeRepository{}
	service := GetInstance(repository, nil)
	input := Input{
		Recipient: "Recipient",
		Reason:    "Reason",
		Meals: []MealVolunteeredRequest{
			{Date: "2026-10-20"},
			{Date: "2026-10-21"},
		},
	}

	response, err := service.Create(context.Background(), input)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	if response.Recipient != input.Recipient || response.Reason != input.Reason ||
		!reflect.DeepEqual(response.Meals, input.Meals) {
		t.Fatalf("Create() response = %#v, want response-shaped values from %#v", response, input)
	}
	if repository.createdParams.ID == "" {
		t.Fatal("Create() did not generate a meal ID")
	}
	if len(repository.createdParams.Meals) != len(input.Meals) {
		t.Fatalf("created volunteered meals = %d, want %d", len(repository.createdParams.Meals), len(input.Meals))
	}
	for index, created := range repository.createdParams.Meals {
		if created.ID == "" || created.Date != input.Meals[index].Date {
			t.Fatalf("created volunteered meal %d = %#v, want generated ID and date", index, created)
		}
	}
}

func TestListReadsCachedMeals(t *testing.T) {
	cached := []Response{{ID: "meal-1", Recipient: "Recipient"}}
	encoded, err := json.Marshal(cached)
	if err != nil {
		t.Fatalf("marshal cached meals: %v", err)
	}
	repository := &fakeRepository{
		meals: []domainmeal.Meal{{ID: "database-meal"}},
	}
	service := GetInstance(repository, &fakeCache{value: encoded, found: true})

	actual, err := service.List(context.Background())
	if err != nil {
		t.Fatalf("List() error = %v", err)
	}
	if !reflect.DeepEqual(actual, cached) {
		t.Fatalf("List() = %#v, want cached %#v", actual, cached)
	}
}

func TestUpdateInvalidatesMealCache(t *testing.T) {
	cacheStore := &fakeCache{}
	repository := &fakeRepository{}
	service := GetInstance(repository, cacheStore)
	input := Input{ID: "meal-1", Recipient: "Updated"}

	if _, err := service.Update(context.Background(), input); err != nil {
		t.Fatalf("Update() error = %v", err)
	}
	if repository.updatedParams.ID != input.ID || repository.updatedParams.Recipient != input.Recipient {
		t.Fatalf("Update() params = %#v, want mapped input", repository.updatedParams)
	}
	if cacheStore.deleteKey != "meals" {
		t.Fatalf("cache delete key = %q, want meals", cacheStore.deleteKey)
	}
}

func TestAssignVolunteerMapsNestedUser(t *testing.T) {
	repository := &fakeRepository{}
	service := GetInstance(repository, nil)
	input := VolunteerAssignmentInput{
		MealID: "meal-1",
		VolunteeredMeal: VolunteeredMealInput{
			ID:          "volunteered-meal-1",
			Description: "Delivered",
			Notes:       "Leave at door",
			User:        UserInput{ID: "user-1"},
		},
	}

	response, err := service.AssignVolunteer(context.Background(), input)
	if err != nil {
		t.Fatalf("AssignVolunteer() error = %v", err)
	}
	if response.MealID != input.MealID ||
		response.VolunteeredMeal.ID != input.VolunteeredMeal.ID ||
		response.VolunteeredMeal.Description != input.VolunteeredMeal.Description ||
		response.VolunteeredMeal.Notes != input.VolunteeredMeal.Notes ||
		response.VolunteeredMeal.User.ID != input.VolunteeredMeal.User.ID {
		t.Fatalf("AssignVolunteer() response = %#v, want mapped response values from %#v", response, input)
	}
	if repository.assignParams.MealID != input.MealID ||
		repository.assignParams.VolunteeredMealID != input.VolunteeredMeal.ID ||
		repository.assignParams.UserID != input.VolunteeredMeal.User.ID {
		t.Fatalf("AssignVolunteer() params = %#v, want nested IDs", repository.assignParams)
	}
}

func TestResponseMapsNestedMealAggregate(t *testing.T) {
	createdOn := time.Date(2026, time.October, 16, 0, 0, 0, 0, time.UTC)
	response := toResponse(domainmeal.Meal{
		ID:        "meal-1",
		CreatedOn: createdOn,
		User:      &domainuser.User{ID: "user-1", CreatedOn: createdOn},
		Meals: []domainmeal.VolunteeredMeal{{
			ID:        "volunteered-1",
			CreatedOn: createdOn,
			User:      &domainuser.User{ID: "user-2"},
			Updates:   []domainmeal.Update{{ID: "update-1", Message: "Updated"}},
		}},
	})
	if response.User.ID != "user-1" || len(response.Meals) != 1 || response.Meals[0].User.ID != "user-2" {
		t.Fatalf("toResponse() = %#v, want nested users and meals", response)
	}
	if response.Meals[0].CreatedOn != createdOn.Format(time.RFC3339Nano) {
		t.Fatalf("volunteered meal CreatedOn = %q, want RFC3339", response.Meals[0].CreatedOn)
	}
}

func TestDeleteMapsNotFound(t *testing.T) {
	service := GetInstance(&fakeRepository{deleteErr: domainmeal.ErrNotFound}, nil)
	_, err := service.Delete(context.Background(), "meal-1")
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("Delete() error = %v, want %v", err, ErrNotFound)
	}
}

var _ cache.Store = (*fakeCache)(nil)

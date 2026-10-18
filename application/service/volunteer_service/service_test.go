package volunteer_service

import (
	"context"
	"encoding/json"
	"errors"
	"reflect"
	"testing"
	"time"

	"Momentum/application/cache"
	domainuser "Momentum/domain/user"
	domainservice "Momentum/domain/volunteer_service"
)

type fakeRepository struct {
	services      []domainservice.Service
	detail        domainservice.ServiceDetail
	createParams  domainservice.CreateParams
	createErr     error
	updatedParams domainservice.UpdateDayParams
	updatedDay    domainservice.VolunteerServiceDay
	updateDayErr  error
	listCalls     int
}

func (repository *fakeRepository) ListServices(context.Context) ([]domainservice.Service, error) {
	repository.listCalls++
	return repository.services, nil
}

func (repository *fakeRepository) GetByType(context.Context, string) (domainservice.ServiceDetail, error) {
	return repository.detail, nil
}

func (repository *fakeRepository) Create(_ context.Context, params domainservice.CreateParams) error {
	repository.createParams = params
	return repository.createErr
}

func (repository *fakeRepository) UpdateDay(_ context.Context, params domainservice.UpdateDayParams) (domainservice.VolunteerServiceDay, error) {
	repository.updatedParams = params
	return repository.updatedDay, repository.updateDayErr
}

type fakeCache struct {
	value   []byte
	found   bool
	getErr  error
	setErr  error
	setKey  string
	setData []byte
	setTTL  time.Duration
}

func (cacheStore *fakeCache) Get(context.Context, string) ([]byte, bool, error) {
	return cacheStore.value, cacheStore.found, cacheStore.getErr
}

func (cacheStore *fakeCache) Set(_ context.Context, key string, value []byte, ttl time.Duration) error {
	cacheStore.setKey = key
	cacheStore.setData = value
	cacheStore.setTTL = ttl
	return cacheStore.setErr
}

func (cacheStore *fakeCache) Delete(context.Context, string) error {
	return nil
}

func TestCreateReturnsExplicitResponse(t *testing.T) {
	repository := &fakeRepository{}
	service := GetInstance(repository, nil)
	input := CreateInput{
		ID:          "request-id",
		Organizer:   "Organizer",
		Description: "Description",
		Email:       "organizer@example.com",
		Phone:       "555-0100",
		Days:        []CreateDayInput{{Date: "2026-04-01"}, {Date: "2026-04-08"}},
		UserID:      "user-1",
		ServiceID:   "service-1",
	}

	response, err := service.Create(context.Background(), input)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	expected := CreateResponse{
		ID:          input.ID,
		Organizer:   input.Organizer,
		Description: input.Description,
		Email:       input.Email,
		Phone:       input.Phone,
		Days:        input.Days,
		UserID:      input.UserID,
		ServiceID:   input.ServiceID,
	}
	if !reflect.DeepEqual(response, expected) {
		t.Fatalf("Create() response = %#v, want %#v", response, expected)
	}
	if repository.createParams.ID == "" {
		t.Fatal("Create() did not generate a persisted volunteer service ID")
	}
	if len(repository.createParams.Days) != len(input.Days) {
		t.Fatalf("persisted days = %d, want %d", len(repository.createParams.Days), len(input.Days))
	}
	for index, day := range repository.createParams.Days {
		if day.ID == "" {
			t.Fatalf("persisted day %d has no generated ID", index)
		}
		if day.Date != input.Days[index].Date {
			t.Fatalf("persisted day %d date = %q, want %q", index, day.Date, input.Days[index].Date)
		}
	}
}

func TestListServicesReadsCachedCatalog(t *testing.T) {
	cached := []ServiceResponse{{ID: "service-1", Name: "Food", Type: "food", Image: "food.png"}}
	encoded, err := json.Marshal(cached)
	if err != nil {
		t.Fatalf("marshal cached catalog: %v", err)
	}

	repository := &fakeRepository{
		services: []domainservice.Service{{ID: "database-service"}},
	}
	cacheStore := &fakeCache{value: encoded, found: true}
	service := GetInstance(repository, cacheStore)

	actual, err := service.ListServices(context.Background())
	if err != nil {
		t.Fatalf("ListServices() error = %v", err)
	}
	if !reflect.DeepEqual(actual, cached) {
		t.Fatalf("ListServices() = %#v, want cached %#v", actual, cached)
	}
	if repository.listCalls != 0 {
		t.Fatalf("ListServices() called repository %d times on cache hit, want 0", repository.listCalls)
	}
}

func TestListServicesCachesRepositoryResult(t *testing.T) {
	repository := &fakeRepository{
		services: []domainservice.Service{{
			ID:    "service-1",
			Name:  "Food",
			Type:  "food",
			Image: "food.png",
		}},
	}
	cacheStore := &fakeCache{}
	service := GetInstance(repository, cacheStore)

	actual, err := service.ListServices(context.Background())
	if err != nil {
		t.Fatalf("ListServices() error = %v", err)
	}
	if len(actual) != 1 || actual[0].ID != "service-1" {
		t.Fatalf("ListServices() = %#v, want catalog response", actual)
	}
	if cacheStore.setKey != "tabs" {
		t.Fatalf("cache key = %q, want %q", cacheStore.setKey, "tabs")
	}
	if cacheStore.setTTL != servicesCacheTTL {
		t.Fatalf("cache TTL = %v, want %v", cacheStore.setTTL, servicesCacheTTL)
	}
}

func TestGetByTypeMapsAggregate(t *testing.T) {
	createdOn := time.Date(2026, time.January, 1, 0, 0, 0, 0, time.UTC)
	repository := &fakeRepository{
		detail: domainservice.ServiceDetail{
			Service: domainservice.Service{ID: "service-1", Name: "Food", Type: "food", Image: "food.png"},
			Services: []domainservice.VolunteerService{{
				ID:          "volunteer-1",
				Description: "Serve food",
				ServiceID:   "service-1",
				Days: []domainservice.VolunteerServiceDay{{
					ID:    "day-1",
					Date:  "2026-04-01",
					User:  &domainuser.User{ID: "user-1", Email: "user@example.com", CreatedOn: createdOn},
					Notes: "Bring supplies",
				}},
			}},
		},
	}
	service := GetInstance(repository, nil)

	actual, err := service.GetByType(context.Background(), "food")
	if err != nil {
		t.Fatalf("GetByType() error = %v", err)
	}
	if actual.ID != "service-1" || len(actual.Services) != 1 || len(actual.Services[0].Days) != 1 {
		t.Fatalf("GetByType() = %#v, want hydrated aggregate", actual)
	}
	if actual.Services[0].Days[0].User.ID != "user-1" {
		t.Fatalf("hydrated day user ID = %q, want %q", actual.Services[0].Days[0].User.ID, "user-1")
	}
	if actual.Services[0].Days[0].User.CreatedOn != createdOn.Format(time.RFC3339Nano) {
		t.Fatalf("hydrated day user CreatedOn = %q, want %q", actual.Services[0].Days[0].User.CreatedOn, createdOn.Format(time.RFC3339Nano))
	}
}

func TestUpdateDayReturnsExplicitResponse(t *testing.T) {
	repository := &fakeRepository{
		updatedDay: domainservice.VolunteerServiceDay{ID: "day-1"},
	}
	service := GetInstance(repository, nil)
	input := UpdateDayInput{
		ID:    "day-1",
		Notes: "Updated notes",
		Date:  "2026-04-01",
		User:  UserResponse{ID: "user-1"},
	}

	response, err := service.UpdateDay(context.Background(), input)
	if err != nil {
		t.Fatalf("UpdateDay() error = %v", err)
	}
	expected := UpdateDayResponse{
		ID:    input.ID,
		Notes: input.Notes,
		Date:  input.Date,
		User:  input.User,
	}
	if !reflect.DeepEqual(response, expected) {
		t.Fatalf("UpdateDay() response = %#v, want %#v", response, expected)
	}
	if repository.updatedParams.ID != input.ID || repository.updatedParams.Notes != input.Notes || repository.updatedParams.UserID != input.User.ID {
		t.Fatalf("UpdateDay() params = %#v, want ID/notes/user ID from input", repository.updatedParams)
	}
}

func TestGetByTypeRejectsEmptyType(t *testing.T) {
	service := GetInstance(&fakeRepository{}, nil)

	_, err := service.GetByType(context.Background(), "")
	if !errors.Is(err, ErrInvalidInput) {
		t.Fatalf("GetByType() error = %v, want %v", err, ErrInvalidInput)
	}
}

var _ cache.Store = (*fakeCache)(nil)

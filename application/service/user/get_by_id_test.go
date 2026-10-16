package user

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	domainuser "Momentum/domain/user"
)

type fakeRepository struct {
	user         domainuser.User
	err          error
	calls        int
	createCalls  int
	createParams domainuser.CreateParams
	updateCalls  int
	updateParams domainuser.UpdateParams
	deleteCalls  int
	deleteID     string
}

func (repository *fakeRepository) GetByID(context.Context, string) (domainuser.User, error) {
	repository.calls++
	return repository.user, repository.err
}

func (repository *fakeRepository) Create(_ context.Context, params domainuser.CreateParams) (domainuser.User, error) {
	repository.createCalls++
	repository.createParams = params
	return repository.user, repository.err
}

func (repository *fakeRepository) Update(_ context.Context, params domainuser.UpdateParams) (domainuser.User, error) {
	repository.updateCalls++
	repository.updateParams = params
	return repository.user, repository.err
}

func (repository *fakeRepository) Delete(_ context.Context, id string) (domainuser.User, error) {
	repository.deleteCalls++
	repository.deleteID = id
	return repository.user, repository.err
}

type fakeCache struct {
	value       []byte
	found       bool
	getErr      error
	setErr      error
	getCalls    int
	setCalls    int
	key         string
	ttl         time.Duration
	deleteCalls int
	deletedKey  string
}

func (cache *fakeCache) Get(context.Context, string) ([]byte, bool, error) {
	cache.getCalls++
	return cache.value, cache.found, cache.getErr
}

func (cache *fakeCache) Set(_ context.Context, key string, value []byte, ttl time.Duration) error {
	cache.setCalls++
	cache.key = key
	cache.value = value
	cache.found = true
	cache.ttl = ttl
	return cache.setErr
}

func (cache *fakeCache) Delete(_ context.Context, key string) error {
	cache.deleteCalls++
	cache.deletedKey = key
	return nil
}

func TestGetByIDReturnsCachedUserWithoutRepositoryCall(t *testing.T) {
	cachedUser := domainuser.User{
		ID:       "user-1",
		Email:    "cached@example.com",
		Fullname: "Cached User",
	}
	cachedValue, err := json.Marshal(cachedUser)
	if err != nil {
		t.Fatalf("marshal cached user: %v", err)
	}

	repository := &fakeRepository{
		user: domainuser.User{ID: "user-1", Email: "database@example.com"},
	}
	cache := &fakeCache{value: cachedValue, found: true}
	service := GetInstance(repository, cache)

	response, err := service.GetByID(context.Background(), "user-1")
	if err != nil {
		t.Fatalf("GetByID() error = %v", err)
	}

	if response.Email != cachedUser.Email {
		t.Fatalf("GetByID() email = %q, want %q", response.Email, cachedUser.Email)
	}
	if repository.calls != 0 {
		t.Fatalf("repository calls = %d, want 0", repository.calls)
	}
	if cache.setCalls != 0 {
		t.Fatalf("cache set calls = %d, want 0", cache.setCalls)
	}
}

func TestGetByIDLoadsAndCachesOnCacheMiss(t *testing.T) {
	repository := &fakeRepository{
		user: domainuser.User{
			ID:       "user-1",
			Email:    "database@example.com",
			Fullname: "Database User",
		},
	}
	cache := &fakeCache{}
	service := GetInstance(repository, cache)

	response, err := service.GetByID(context.Background(), "user-1")
	if err != nil {
		t.Fatalf("GetByID() error = %v", err)
	}

	if response.Email != repository.user.Email {
		t.Fatalf("GetByID() email = %q, want %q", response.Email, repository.user.Email)
	}
	if repository.calls != 1 {
		t.Fatalf("repository calls = %d, want 1", repository.calls)
	}
	if cache.setCalls != 1 {
		t.Fatalf("cache set calls = %d, want 1", cache.setCalls)
	}
	if cache.key != userCacheKeyPrefix+"user-1" {
		t.Fatalf("cache key = %q, want %q", cache.key, userCacheKeyPrefix+"user-1")
	}
	if cache.ttl != userCacheTTL {
		t.Fatalf("cache ttl = %s, want %s", cache.ttl, userCacheTTL)
	}
}

func TestGetByIDFallsBackWhenCacheFails(t *testing.T) {
	repository := &fakeRepository{
		user: domainuser.User{ID: "user-1", Email: "database@example.com"},
	}
	cache := &fakeCache{getErr: errors.New("redis unavailable")}
	service := GetInstance(repository, cache)

	response, err := service.GetByID(context.Background(), "user-1")
	if err != nil {
		t.Fatalf("GetByID() error = %v", err)
	}

	if response.Email != repository.user.Email {
		t.Fatalf("GetByID() email = %q, want %q", response.Email, repository.user.Email)
	}
	if repository.calls != 1 {
		t.Fatalf("repository calls = %d, want 1", repository.calls)
	}
}

func TestGetByIDTranslatesMissingUser(t *testing.T) {
	repository := &fakeRepository{err: domainuser.ErrNotFound}
	service := GetInstance(repository, nil)

	_, err := service.GetByID(context.Background(), "missing-user")
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("GetByID() error = %v, want %v", err, ErrNotFound)
	}
}

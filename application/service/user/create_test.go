package user

import (
	"context"
	"testing"

	domainuser "Momentum/domain/user"
)

func TestCreateReturnsResponseAndPopulatesCache(t *testing.T) {
	repository := &fakeRepository{
		user: domainuser.User{
			ID:       "user-1",
			Email:    "created@example.com",
			Fullname: "Created User",
			Phone:    "555-0100",
		},
	}
	cache := &fakeCache{}
	service := GetInstance(repository, cache)

	response, err := service.Create(context.Background(), CreateInput{
		ID:       "user-1",
		Email:    "created@example.com",
		Fullname: "Created User",
		Phone:    "555-0100",
	})
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	if response.ID != repository.user.ID {
		t.Fatalf("Create() id = %q, want %q", response.ID, repository.user.ID)
	}
	if repository.createCalls != 1 {
		t.Fatalf("repository create calls = %d, want 1", repository.createCalls)
	}
	if repository.createParams.Fullname != "Created User" {
		t.Fatalf("repository fullname = %q, want %q", repository.createParams.Fullname, "Created User")
	}
	if cache.setCalls != 1 {
		t.Fatalf("cache set calls = %d, want 1", cache.setCalls)
	}
	if cache.key != userCacheKeyPrefix+"user-1" {
		t.Fatalf("cache key = %q, want %q", cache.key, userCacheKeyPrefix+"user-1")
	}
}

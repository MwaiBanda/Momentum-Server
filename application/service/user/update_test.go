package user

import (
	"context"
	"testing"

	domainuser "Momentum/domain/user"
)

func TestUpdateReturnsResponseAndRefreshesCache(t *testing.T) {
	repository := &fakeRepository{
		user: domainuser.User{
			ID:       "user-1",
			Email:    "updated@example.com",
			Fullname: "Updated User",
			Phone:    "555-0101",
		},
	}
	cache := &fakeCache{}
	service := GetInstance(repository, cache)

	response, err := service.Update(context.Background(), UpdateInput{
		ID:       "user-1",
		Email:    "updated@example.com",
		Fullname: "Updated User",
		Phone:    "555-0101",
	})
	if err != nil {
		t.Fatalf("Update() error = %v", err)
	}

	if response.Email != repository.user.Email {
		t.Fatalf("Update() email = %q, want %q", response.Email, repository.user.Email)
	}
	if repository.updateCalls != 1 {
		t.Fatalf("repository update calls = %d, want 1", repository.updateCalls)
	}
	if repository.updateParams.Fullname != "Updated User" {
		t.Fatalf("repository fullname = %q, want %q", repository.updateParams.Fullname, "Updated User")
	}
	if cache.setCalls != 1 {
		t.Fatalf("cache set calls = %d, want 1", cache.setCalls)
	}
	if cache.key != userCacheKeyPrefix+"user-1" {
		t.Fatalf("cache key = %q, want %q", cache.key, userCacheKeyPrefix+"user-1")
	}
}

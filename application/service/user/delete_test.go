package user

import (
	"context"
	"errors"
	"testing"

	domainuser "Momentum/domain/user"
)

func TestDeleteReturnsResponseAndInvalidatesCache(t *testing.T) {
	repository := &fakeRepository{
		user: domainuser.User{
			ID:       "user-1",
			Email:    "deleted@example.com",
			Fullname: "Deleted User",
		},
	}
	cache := &fakeCache{}
	service := GetInstance(repository, cache)

	response, err := service.Delete(context.Background(), "user-1")
	if err != nil {
		t.Fatalf("Delete() error = %v", err)
	}

	if response.ID != repository.user.ID {
		t.Fatalf("Delete() id = %q, want %q", response.ID, repository.user.ID)
	}
	if repository.deleteCalls != 1 {
		t.Fatalf("repository delete calls = %d, want 1", repository.deleteCalls)
	}
	if repository.deleteID != "user-1" {
		t.Fatalf("repository delete id = %q, want %q", repository.deleteID, "user-1")
	}
	if cache.deleteCalls != 1 {
		t.Fatalf("cache delete calls = %d, want 1", cache.deleteCalls)
	}
	if cache.deletedKey != userCacheKeyPrefix+"user-1" {
		t.Fatalf("cache delete key = %q, want %q", cache.deletedKey, userCacheKeyPrefix+"user-1")
	}
}

func TestDeleteTranslatesMissingUser(t *testing.T) {
	repository := &fakeRepository{err: domainuser.ErrNotFound}
	service := GetInstance(repository, nil)

	_, err := service.Delete(context.Background(), "missing-user")
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("Delete() error = %v, want %v", err, ErrNotFound)
	}
}

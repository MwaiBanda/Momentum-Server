package cache

import (
	"context"
	"errors"
	"testing"
	"time"
)

type fakeStore struct {
	value   []byte
	found   bool
	getErr  error
	setErr  error
	setKey  string
	setTTL  time.Duration
	setData []byte
}

func (store *fakeStore) Get(context.Context, string) ([]byte, bool, error) {
	return store.value, store.found, store.getErr
}

func (store *fakeStore) Set(_ context.Context, key string, value []byte, ttl time.Duration) error {
	store.setKey = key
	store.setData = value
	store.setTTL = ttl
	return store.setErr
}

func (store *fakeStore) Delete(context.Context, string) error {
	return nil
}

type cachedUser struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func TestGetDecodesJSON(t *testing.T) {
	store := &fakeStore{
		value: []byte(`{"id":"user-1","email":"user@example.com"}`),
		found: true,
	}

	value, found, err := Get[cachedUser](context.Background(), store, "user-1")
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	if !found {
		t.Fatal("Get() found = false, want true")
	}
	if value.ID != "user-1" || value.Email != "user@example.com" {
		t.Fatalf("Get() value = %#v, want decoded user", value)
	}
}

func TestGetReturnsMissWithoutDecoding(t *testing.T) {
	store := &fakeStore{value: []byte(`{invalid`)}

	value, found, err := Get[cachedUser](context.Background(), store, "missing")
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	if found {
		t.Fatal("Get() found = true, want false")
	}
	if value != nil {
		t.Fatalf("Get() value = %#v, want nil", value)
	}
}

func TestGetReturnsInvalidJSONError(t *testing.T) {
	store := &fakeStore{
		value: []byte(`{invalid`),
		found: true,
	}

	value, found, err := Get[cachedUser](context.Background(), store, "invalid")
	if err == nil {
		t.Fatal("Get() error = nil, want JSON decoding error")
	}
	if found {
		t.Fatal("Get() found = true, want false for invalid payload")
	}
	if value != nil {
		t.Fatalf("Get() value = %#v, want nil", value)
	}
}

func TestSetEncodesJSONAndPreservesMetadata(t *testing.T) {
	store := &fakeStore{}
	ttl := 15 * time.Minute

	if err := Set(context.Background(), store, "user-1", cachedUser{
		ID:    "user-1",
		Email: "user@example.com",
	}, ttl); err != nil {
		t.Fatalf("Set() error = %v", err)
	}
	if store.setKey != "user-1" {
		t.Fatalf("Set() key = %q, want %q", store.setKey, "user-1")
	}
	if string(store.setData) != `{"id":"user-1","email":"user@example.com"}` {
		t.Fatalf("Set() data = %q, want encoded JSON", store.setData)
	}
	if store.setTTL != ttl {
		t.Fatalf("Set() TTL = %v, want %v", store.setTTL, ttl)
	}
}

func TestSetPropagatesStoreError(t *testing.T) {
	expectedErr := errors.New("cache unavailable")
	store := &fakeStore{setErr: expectedErr}

	err := Set(context.Background(), store, "user-1", cachedUser{ID: "user-1"}, time.Minute)
	if !errors.Is(err, expectedErr) {
		t.Fatalf("Set() error = %v, want %v", err, expectedErr)
	}
}

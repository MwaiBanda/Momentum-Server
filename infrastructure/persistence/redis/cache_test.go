package redis

import (
	"context"
	"errors"
	"testing"
	"time"

	applicationcache "Momentum/application/cache"

	"github.com/alicebob/miniredis/v2"
	"github.com/redis/go-redis/v9"
)

func newTestCache(t *testing.T) (*Cache, *miniredis.Miniredis) {
	t.Helper()

	server := miniredis.RunT(t)
	client := redis.NewClient(&redis.Options{Addr: server.Addr()})
	t.Cleanup(func() {
		_ = client.Close()
	})

	return NewCache(client), server
}

func TestCacheGetReturnsMissForMissingKey(t *testing.T) {
	cache, _ := newTestCache(t)

	value, found, err := cache.Get(context.Background(), "missing")
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	if found {
		t.Fatal("Get() found = true, want false")
	}
	if value != nil {
		t.Fatalf("Get() value = %q, want nil", value)
	}
}

func TestCacheSetGetAndDelete(t *testing.T) {
	cache, server := newTestCache(t)
	key := "user-1"
	expected := []byte(`{"id":"user-1"}`)

	if err := cache.Set(context.Background(), key, expected, time.Hour); err != nil {
		t.Fatalf("Set() error = %v", err)
	}
	if !server.Exists(key) {
		t.Fatal("Set() did not create the Redis key")
	}

	actual, found, err := cache.Get(context.Background(), key)
	if err != nil {
		t.Fatalf("Get() error = %v", err)
	}
	if !found {
		t.Fatal("Get() found = false, want true")
	}
	if string(actual) != string(expected) {
		t.Fatalf("Get() value = %q, want %q", actual, expected)
	}

	if err := cache.Delete(context.Background(), key); err != nil {
		t.Fatalf("Delete() error = %v", err)
	}
	if server.Exists(key) {
		t.Fatal("Delete() did not remove the Redis key")
	}
}

func TestCacheSetExpiresAfterTTL(t *testing.T) {
	cache, server := newTestCache(t)
	key := "expiring-user"

	if err := cache.Set(context.Background(), key, []byte("value"), time.Minute); err != nil {
		t.Fatalf("Set() error = %v", err)
	}

	server.FastForward(time.Minute)

	value, found, err := cache.Get(context.Background(), key)
	if err != nil {
		t.Fatalf("Get() after expiration error = %v", err)
	}
	if found {
		t.Fatal("Get() found = true after expiration, want false")
	}
	if value != nil {
		t.Fatalf("Get() value after expiration = %q, want nil", value)
	}
}

func TestCacheReturnsUnavailableWhenClientIsMissing(t *testing.T) {
	cache := NewCache(nil)

	_, found, err := cache.Get(context.Background(), "key")
	if !errors.Is(err, applicationcache.ErrUnavailable) {
		t.Fatalf("Get() error = %v, want %v", err, applicationcache.ErrUnavailable)
	}
	if found {
		t.Fatal("Get() found = true, want false")
	}

	if err := cache.Set(context.Background(), "key", []byte("value"), time.Minute); !errors.Is(err, applicationcache.ErrUnavailable) {
		t.Fatalf("Set() error = %v, want %v", err, applicationcache.ErrUnavailable)
	}
	if err := cache.Delete(context.Background(), "key"); !errors.Is(err, applicationcache.ErrUnavailable) {
		t.Fatalf("Delete() error = %v, want %v", err, applicationcache.ErrUnavailable)
	}
}

package cache

import (
	"context"
	"encoding/json"
	"errors"
	"time"
)

var ErrUnavailable = errors.New("cache unavailable")

type Store interface {
	Get(ctx context.Context, key string) ([]byte, bool, error)
	Set(ctx context.Context, key string, value []byte, ttl time.Duration) error
	Delete(ctx context.Context, key string) error
}

type CounterStore interface {
	Increment(ctx context.Context, key string) (int64, error)
}

func Get[T any](ctx context.Context, store Store, key string) (*T, bool, error) {
	value, found, err := store.Get(ctx, key)
	if err != nil || !found {
		return nil, found, err
	}

	var decoded T
	if err := json.Unmarshal(value, &decoded); err != nil {
		return nil, false, err
	}

	return &decoded, true, nil
}

func Set[T any](ctx context.Context, store Store, key string, value T, ttl time.Duration) error {
	encoded, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return store.Set(ctx, key, encoded, ttl)
}

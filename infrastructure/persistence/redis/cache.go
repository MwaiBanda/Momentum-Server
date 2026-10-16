package redis

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	client redis.UniversalClient
}

func NewCache(client redis.UniversalClient) *Cache {
	return &Cache{client: client}
}

func (cache *Cache) Get(ctx context.Context, key string) ([]byte, bool, error) {
	if cache == nil || cache.client == nil {
		return nil, false, errors.New("redis cache unavailable")
	}

	value, err := cache.client.Get(ctx, key).Bytes()
	if errors.Is(err, redis.Nil) {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, err
	}

	return value, true, nil
}

func (cache *Cache) Set(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	if cache == nil || cache.client == nil {
		return errors.New("redis cache unavailable")
	}

	return cache.client.Set(ctx, key, value, ttl).Err()
}

func (cache *Cache) Delete(ctx context.Context, key string) error {
	if cache == nil || cache.client == nil {
		return errors.New("redis cache unavailable")
	}

	return cache.client.Del(ctx, key).Err()
}

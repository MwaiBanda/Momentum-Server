package redis

import (
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

func Open(redisURL string) (*redis.Client, error) {
	options, err := redis.ParseURL(redisURL)
	if err != nil {
		return nil, fmt.Errorf("parse Redis URL: %w", err)
	}

	client := redis.NewClient(options)
	log.Println("Connected to Redis")
	return client, nil
}

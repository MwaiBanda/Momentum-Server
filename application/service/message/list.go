package message

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"Momentum/application/cache"
	"Momentum/application/telemetry"
)

const (
	messageCacheTTL        = 24 * time.Hour
	messageCacheVersionKey = "messages:version"
)

func (service *Service) ListPublishedByUser(ctx context.Context, userID string) (MessageListResponse, error) {
	ctx, span := telemetry.Start(ctx, "application.message.list_published_by_user")
	defer span.End()

	if userID == "" {
		return MessageListResponse{}, ErrInvalidInput
	}

	cacheKey, cacheEnabled := service.versionedMessageCacheKey(ctx, userID)
	if userID != adminUserID && cacheEnabled {
		cached, found, err := cache.Get[[]MessageResponse](ctx, service.cacheStore, cacheKey)
		if err != nil {
			log.Printf("[message-cache] get %q: %v", cacheKey, err)
		} else if found {
			return MessageListResponse{Data: *cached}, nil
		}
	}

	messages, err := service.repository.ListPublishedByUser(ctx, userID)
	if err != nil {
		return MessageListResponse{}, err
	}

	response := toMessageListResponse(messages)
	if userID != adminUserID && cacheEnabled {
		if err := cache.Set(ctx, service.cacheStore, cacheKey, response.Data, messageCacheTTL); err != nil {
			log.Printf("[message-cache] set %q: %v", cacheKey, err)
		}
	}

	return response, nil
}

func (service *Service) ListUnpublished(ctx context.Context) (MessageListResponse, error) {
	ctx, span := telemetry.Start(ctx, "application.message.list_unpublished")
	defer span.End()

	messages, err := service.repository.ListUnpublished(ctx)
	if err != nil {
		return MessageListResponse{}, err
	}
	return toMessageListResponse(messages), nil
}

func (service *Service) versionedMessageCacheKey(ctx context.Context, userID string) (string, bool) {
	if service.cacheStore == nil || userID == adminUserID {
		return "", false
	}

	version, err := service.messageCacheVersion(ctx)
	if err != nil {
		log.Printf("[message-cache] read version: %v", err)
		return "", false
	}

	return fmt.Sprintf("%s-v%d-%s", messageCacheKey, version, userID), true
}

func (service *Service) messageCacheVersion(ctx context.Context) (int64, error) {
	if _, ok := service.cacheStore.(cache.CounterStore); !ok {
		return 0, fmt.Errorf("cache store does not support atomic counters")
	}

	value, found, err := service.cacheStore.Get(ctx, messageCacheVersionKey)
	if err != nil {
		return 0, err
	}
	if !found {
		return 0, nil
	}

	version, err := strconv.ParseInt(string(value), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("parse cache version %q: %w", value, err)
	}
	return version, nil
}

func (service *Service) bumpMessageCacheVersion(ctx context.Context) {
	if service.cacheStore == nil {
		return
	}

	counterStore, ok := service.cacheStore.(cache.CounterStore)
	if !ok {
		log.Printf("[message-cache] cache store does not support atomic versioning")
		return
	}

	if _, err := counterStore.Increment(ctx, messageCacheVersionKey); err != nil {
		log.Printf("[message-cache] increment version: %v", err)
	}
}

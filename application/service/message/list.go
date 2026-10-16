package message

import (
	"context"
	"log"
	"time"

	"Momentum/application/cache"
	"Momentum/constants"
)

const messageCacheTTL = 24 * time.Hour

func (service *Service) ListPublishedByUser(ctx context.Context, userID string) (MessageListResponse, error) {
	if userID == "" {
		return MessageListResponse{}, ErrInvalidInput
	}

	if userID != constants.Admin && service.cacheStore != nil {
		cached, found, err := cache.Get[[]MessageResponse](ctx, service.cacheStore, messageCacheKey(userID))
		if err != nil {
			log.Printf("[message-cache] get %q: %v", messageCacheKey(userID), err)
		} else if found {
			return MessageListResponse{Data: *cached}, nil
		}
	}

	messages, err := service.repository.ListPublishedByUser(ctx, userID)
	if err != nil {
		return MessageListResponse{}, err
	}

	response := toMessageListResponse(messages)
	if userID != constants.Admin {
		service.cacheMessages(ctx, userID, response.Data)
	}

	return response, nil
}

func (service *Service) ListUnpublished(ctx context.Context) (MessageListResponse, error) {
	messages, err := service.repository.ListUnpublished(ctx)
	if err != nil {
		return MessageListResponse{}, err
	}
	return toMessageListResponse(messages), nil
}

func messageCacheKey(userID string) string {
	return constants.MessageKey + "-" + userID
}

func (service *Service) cacheMessages(ctx context.Context, userID string, messages []MessageResponse) {
	if service.cacheStore == nil {
		return
	}

	key := messageCacheKey(userID)
	if err := cache.Set(ctx, service.cacheStore, key, messages, messageCacheTTL); err != nil {
		log.Printf("[message-cache] set %q: %v", key, err)
		return
	}

	service.addCacheSession(ctx, key)
}

func (service *Service) invalidateMessageCache(ctx context.Context, userID string) {
	if service.cacheStore == nil || userID == "" {
		return
	}

	key := messageCacheKey(userID)
	if err := service.cacheStore.Delete(ctx, key); err != nil {
		log.Printf("[message-cache] delete %q: %v", key, err)
	}
}

func (service *Service) invalidateAllMessageCaches(ctx context.Context) {
	if service.cacheStore == nil {
		return
	}

	keys, found, err := cache.Get[[]string](ctx, service.cacheStore, constants.UserCachedSessionKey)
	if err != nil {
		log.Printf("[message-cache] get sessions: %v", err)
		return
	}
	if !found {
		return
	}

	for _, key := range *keys {
		if err := service.cacheStore.Delete(ctx, key); err != nil {
			log.Printf("[message-cache] delete %q: %v", key, err)
		}
	}
	if err := service.cacheStore.Delete(ctx, constants.UserCachedSessionKey); err != nil {
		log.Printf("[message-cache] delete sessions: %v", err)
	}
}

func (service *Service) addCacheSession(ctx context.Context, key string) {
	if service.cacheStore == nil {
		return
	}

	sessions, found, err := cache.Get[[]string](ctx, service.cacheStore, constants.UserCachedSessionKey)
	if err != nil {
		log.Printf("[message-cache] get sessions: %v", err)
		return
	}

	if !found {
		sessions = &[]string{key}
	} else {
		for _, existing := range *sessions {
			if existing == key {
				return
			}
		}
		*sessions = append(*sessions, key)
	}

	if err := cache.Set(ctx, service.cacheStore, constants.UserCachedSessionKey, *sessions, messageCacheTTL); err != nil {
		log.Printf("[message-cache] set sessions: %v", err)
	}
}

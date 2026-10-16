package user

import (
	"context"
	"log"
	"time"

	"Momentum/application/cache"
	domainuser "Momentum/domain/user"
)

const (
	userCacheKeyPrefix = "user-"
	userCacheTTL       = 24 * time.Hour
)

func (service *Service) cacheUser(ctx context.Context, userEntity domainuser.User) {
	if service.cacheStore == nil {
		return
	}

	if err := cache.Set(ctx, service.cacheStore, userCacheKeyPrefix+userEntity.ID, userEntity, userCacheTTL); err != nil {
		log.Printf("[user-cache] set %q: %v", userEntity.ID, err)
	}
}

func (service *Service) deleteCachedUser(ctx context.Context, id string) {
	if service.cacheStore == nil {
		return
	}

	if err := service.cacheStore.Delete(ctx, userCacheKeyPrefix+id); err != nil {
		log.Printf("[user-cache] delete %q: %v", id, err)
	}
}

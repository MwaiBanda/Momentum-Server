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

package user

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"Momentum/application/cache"
	domainuser "Momentum/domain/user"
)

const (
	userCacheKeyPrefix = "user-"
	userCacheTTL       = 24 * time.Hour
)

func (service *Service) GetByID(ctx context.Context, id string) (Response, error) {
	if service.cacheStore != nil {
		cachedUser, found, err := cache.Get[domainuser.User](ctx, service.cacheStore, userCacheKeyPrefix+id)
		if err != nil {
			log.Printf("[user-cache] get %q: %v", id, err)
		} else if found {
			return toResponse(*cachedUser), nil
		}
	}

	userEntity, err := service.repository.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Response{}, ErrNotFound
		}
		return Response{}, err
	}

	if service.cacheStore != nil {
		if err := cache.Set(ctx, service.cacheStore, userCacheKeyPrefix+id, userEntity, userCacheTTL); err != nil {
			log.Printf("[user-cache] set %q: %v", id, err)
		}
	}

	return toResponse(userEntity), nil
}

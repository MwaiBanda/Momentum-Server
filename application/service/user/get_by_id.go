package user

import (
	"context"
	"errors"
	"log"

	"Momentum/application/cache"
	domainuser "Momentum/domain/user"
)

func (service *Service) GetByID(ctx context.Context, id string) (Response, error) {
	if service.cacheStore != nil {
		cachedUser, found, err := cache.Get[domainuser.User](ctx, service.cacheStore, userCacheKeyPrefix+id)
		if err != nil {
			log.Printf("[user-cache] get %q: %v", id, err)
			return service.loadByID(ctx, id)
		} else if found {
			return toResponse(*cachedUser), nil
		}
	}

	return service.loadByID(ctx, id)
}

func (service *Service) loadByID(ctx context.Context, id string) (Response, error) {
	userEntity, err := service.repository.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, domainuser.ErrNotFound) {
			return Response{}, ErrNotFound
		}
		return Response{}, err
	}

	service.cacheUser(ctx, userEntity)

	return toResponse(userEntity), nil
}

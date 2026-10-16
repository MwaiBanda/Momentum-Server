package meal

import (
	"context"
	"log"
	"time"

	"Momentum/application/cache"
	"Momentum/constants"
)

const mealsCacheTTL = 24 * time.Hour

func (service *Service) List(ctx context.Context) ([]Response, error) {
	if service == nil || service.repository == nil {
		return nil, ErrInvalidInput
	}

	if service.cacheStore != nil {
		cached, found, err := cache.Get[[]Response](ctx, service.cacheStore, constants.MealsKey)
		if err != nil {
			log.Printf("[meal-cache] get %q: %v", constants.MealsKey, err)
		} else if found {
			return *cached, nil
		}
	}

	meals, err := service.repository.List(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]Response, 0, len(meals))
	for _, entity := range meals {
		responses = append(responses, toResponse(entity))
	}

	if service.cacheStore != nil {
		if err := cache.Set(ctx, service.cacheStore, constants.MealsKey, responses, mealsCacheTTL); err != nil {
			log.Printf("[meal-cache] set %q: %v", constants.MealsKey, err)
		}
	}

	return responses, nil
}

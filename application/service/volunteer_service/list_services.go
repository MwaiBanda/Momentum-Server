package volunteer_service

import (
	"context"
	"log"
	"time"

	"Momentum/application/cache"
	"Momentum/constants"
)

const servicesCacheTTL = 24 * time.Hour

func (service *Service) ListServices(ctx context.Context) ([]ServiceResponse, error) {
	if service.cacheStore != nil {
		cached, found, err := cache.Get[[]ServiceResponse](ctx, service.cacheStore, constants.TabsKey)
		if err != nil {
			log.Printf("[volunteer-service-cache] get %q: %v", constants.TabsKey, err)
		} else if found {
			return *cached, nil
		}
	}

	services, err := service.repository.ListServices(ctx)
	if err != nil {
		return nil, err
	}

	responses := make([]ServiceResponse, 0, len(services))
	for _, catalogService := range services {
		responses = append(responses, toServiceResponse(catalogService))
	}

	if service.cacheStore != nil {
		if err := cache.Set(ctx, service.cacheStore, constants.TabsKey, responses, servicesCacheTTL); err != nil {
			log.Printf("[volunteer-service-cache] set %q: %v", constants.TabsKey, err)
		}
	}

	return responses, nil
}

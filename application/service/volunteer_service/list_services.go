package volunteer_service

import (
	"context"
	"log"
	"time"

	"Momentum/application/cache"
	"Momentum/application/telemetry"
)

const (
	servicesCacheKey = "tabs"
	servicesCacheTTL = 24 * time.Hour
)

func (service *Service) ListServices(ctx context.Context) ([]ServiceResponse, error) {
	ctx, span := telemetry.Start(ctx, "application.volunteer_service.list_services")
	defer span.End()

	if service.cacheStore != nil {
		cached, found, err := cache.Get[[]ServiceResponse](ctx, service.cacheStore, servicesCacheKey)
		if err != nil {
			log.Printf("[volunteer-service-cache] get %q: %v", servicesCacheKey, err)
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
		if err := cache.Set(ctx, service.cacheStore, servicesCacheKey, responses, servicesCacheTTL); err != nil {
			log.Printf("[volunteer-service-cache] set %q: %v", servicesCacheKey, err)
		}
	}

	return responses, nil
}

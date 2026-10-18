package event

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"Momentum/application/cache"
	"Momentum/application/telemetry"
)

const (
	cacheTTL       = 2 * time.Hour
	eventsCacheKey = "events"
)

func (service *Service) List(ctx context.Context) (Response, error) {
	ctx, span := telemetry.Start(ctx, "application.event.list")
	defer span.End()

	if service == nil {
		return Response{}, ErrUnavailable
	}

	if service.cacheStore != nil {
		cached, found, err := cache.Get[Response](ctx, service.cacheStore, eventsCacheKey)
		if err != nil {
			log.Printf("[event-cache] get %q: %v", eventsCacheKey, err)
		} else if found {
			return *cached, nil
		}
	}

	if service.provider == nil {
		return Response{}, ErrUnavailable
	}

	body, err := service.provider.Fetch(ctx)
	if err != nil {
		return Response{}, err
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		return Response{}, fmt.Errorf("decode event response: %w", err)
	}

	if service.cacheStore != nil {
		if err := cache.Set(ctx, service.cacheStore, eventsCacheKey, response, cacheTTL); err != nil {
			log.Printf("[event-cache] set %q: %v", eventsCacheKey, err)
		}
	}

	return response, nil
}

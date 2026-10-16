package event

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"Momentum/application/cache"
	"Momentum/constants"
)

const cacheTTL = 2 * time.Hour

func (service *Service) List(ctx context.Context) (Response, error) {
	if service == nil {
		return Response{}, ErrUnavailable
	}

	if service.cacheStore != nil {
		cached, found, err := cache.Get[Response](ctx, service.cacheStore, constants.EventsKey)
		if err != nil {
			log.Printf("[event-cache] get %q: %v", constants.EventsKey, err)
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
		if err := cache.Set(ctx, service.cacheStore, constants.EventsKey, response, cacheTTL); err != nil {
			log.Printf("[event-cache] set %q: %v", constants.EventsKey, err)
		}
	}

	return response, nil
}

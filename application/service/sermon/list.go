package sermon

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"Momentum/application/cache"
	"Momentum/constants"
)

const cacheTTL = 6 * time.Hour

func (service *Service) List(ctx context.Context, page string) (Response, error) {
	if service == nil {
		return Response{}, ErrUnavailable
	}

	page, err := normalizePage(page)
	if err != nil {
		return Response{}, err
	}

	cacheKey := constants.SermonKeyPrefix + page
	if service.cacheStore != nil {
		cached, found, err := cache.Get[Response](ctx, service.cacheStore, cacheKey)
		if err != nil {
			log.Printf("[sermon-cache] get %q: %v", cacheKey, err)
		} else if found {
			return *cached, nil
		}
	}

	if service.provider == nil {
		return Response{}, ErrUnavailable
	}

	body, err := service.provider.Fetch(ctx, page)
	if err != nil {
		return Response{}, err
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		return Response{}, fmt.Errorf("decode sermon response: %w", err)
	}

	if service.cacheStore != nil {
		if err := cache.Set(ctx, service.cacheStore, cacheKey, response, cacheTTL); err != nil {
			log.Printf("[sermon-cache] set %q: %v", cacheKey, err)
		}
	}

	return response, nil
}

func normalizePage(page string) (string, error) {
	page = strings.TrimSpace(page)
	if page == "" {
		page = "1"
	}

	pageNumber, err := strconv.Atoi(page)
	if err != nil || pageNumber < 1 {
		return "", ErrInvalidPage
	}
	return strconv.Itoa(pageNumber), nil
}

package meal

import (
	"context"
	"errors"

	domainmeal "Momentum/domain/meal"
)

func (service *Service) Delete(ctx context.Context, id string) (Response, error) {
	if service == nil || service.repository == nil || id == "" {
		return Response{}, ErrInvalidInput
	}

	entity, err := service.repository.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, domainmeal.ErrNotFound) {
			return Response{}, ErrNotFound
		}
		return Response{}, err
	}

	if err := service.invalidateCache(ctx); err != nil {
		return Response{}, err
	}
	return toResponse(entity), nil
}

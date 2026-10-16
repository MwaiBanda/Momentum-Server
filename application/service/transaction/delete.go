package transaction

import (
	"context"
	"errors"

	domaintransaction "Momentum/domain/transaction"
)

func (service *Service) Delete(ctx context.Context, id string) (Response, error) {
	if service == nil || service.repository == nil || id == "" {
		return Response{}, ErrInvalidInput
	}

	deleted, err := service.repository.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, domaintransaction.ErrNotFound) {
			return Response{}, ErrNotFound
		}
		return Response{}, err
	}
	return toResponse(deleted), nil
}

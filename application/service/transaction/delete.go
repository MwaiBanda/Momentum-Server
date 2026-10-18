package transaction

import (
	"context"
	"errors"

	"Momentum/application/telemetry"
	domaintransaction "Momentum/domain/transaction"
)

func (service *Service) Delete(ctx context.Context, id string) (Response, error) {
	ctx, span := telemetry.Start(ctx, "application.transaction.delete")
	defer span.End()

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

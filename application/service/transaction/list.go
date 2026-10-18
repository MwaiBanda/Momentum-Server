package transaction

import (
	"context"

	"Momentum/application/telemetry"
	domaintransaction "Momentum/domain/transaction"
)

func (service *Service) List(ctx context.Context) ([]Response, error) {
	ctx, span := telemetry.Start(ctx, "application.transaction.list")
	defer span.End()

	if service == nil || service.repository == nil {
		return nil, ErrInvalidInput
	}

	transactions, err := service.repository.List(ctx)
	if err != nil {
		return nil, err
	}
	return mapResponses(transactions), nil
}

func (service *Service) ListByUser(ctx context.Context, userID string) ([]Response, error) {
	ctx, span := telemetry.Start(ctx, "application.transaction.list_by_user")
	defer span.End()

	if service == nil || service.repository == nil || userID == "" {
		return nil, ErrInvalidInput
	}

	transactions, err := service.repository.ListByUser(ctx, userID)
	if err != nil {
		return nil, err
	}
	return mapResponses(transactions), nil
}

func mapResponses(transactions []domaintransaction.Transaction) []Response {
	responses := make([]Response, 0, len(transactions))
	for _, transaction := range transactions {
		responses = append(responses, toResponse(transaction))
	}
	return responses
}

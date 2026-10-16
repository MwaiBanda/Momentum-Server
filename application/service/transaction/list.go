package transaction

import (
	"context"

	domaintransaction "Momentum/domain/transaction"
)

func (service *Service) List(ctx context.Context) ([]Response, error) {
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

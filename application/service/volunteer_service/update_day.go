package volunteer_service

import (
	"context"
	"errors"

	domainservice "Momentum/domain/volunteer_service"
)

func (service *Service) UpdateDay(ctx context.Context, input UpdateDayInput) (UpdateDayResponse, error) {
	if input.ID == "" {
		return UpdateDayResponse{}, ErrInvalidInput
	}

	_, err := service.repository.UpdateDay(ctx, domainservice.UpdateDayParams{
		ID:     input.ID,
		Notes:  input.Notes,
		UserID: input.User.ID,
	})
	if err != nil {
		if errors.Is(err, domainservice.ErrDayNotFound) {
			return UpdateDayResponse{}, ErrDayNotFound
		}
		return UpdateDayResponse{}, err
	}

	return toUpdateDayResponse(input), nil
}

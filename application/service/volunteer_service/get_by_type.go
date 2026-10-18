package volunteer_service

import (
	"context"
	"errors"

	"Momentum/application/telemetry"
	domainservice "Momentum/domain/volunteer_service"
)

func (service *Service) GetByType(ctx context.Context, serviceType string) (ServiceDetailResponse, error) {
	ctx, span := telemetry.Start(ctx, "application.volunteer_service.get_by_type")
	defer span.End()

	if serviceType == "" {
		return ServiceDetailResponse{}, ErrInvalidInput
	}

	detail, err := service.repository.GetByType(ctx, serviceType)
	if err != nil {
		if errors.Is(err, domainservice.ErrServiceNotFound) {
			return ServiceDetailResponse{}, ErrServiceNotFound
		}
		return ServiceDetailResponse{}, err
	}

	return toDetailResponse(detail), nil
}

package volunteer_service

import (
	"context"

	"Momentum/application/telemetry"
	domainservice "Momentum/domain/volunteer_service"
	"github.com/lucsky/cuid"
)

func (service *Service) Create(ctx context.Context, input CreateInput) (CreateResponse, error) {
	ctx, span := telemetry.Start(ctx, "application.volunteer_service.create")
	defer span.End()

	days := make([]domainservice.CreateDayParams, 0, len(input.Days))
	for _, day := range input.Days {
		days = append(days, domainservice.CreateDayParams{
			ID:   cuid.New(),
			Date: day.Date,
		})
	}

	if err := service.repository.Create(ctx, domainservice.CreateParams{
		ID:          cuid.New(),
		Description: input.Description,
		Email:       input.Email,
		Phone:       input.Phone,
		Organizer:   input.Organizer,
		ServiceID:   input.ServiceID,
		UserID:      input.UserID,
		Days:        days,
	}); err != nil {
		return CreateResponse{}, err
	}

	return toCreateResponse(input), nil
}

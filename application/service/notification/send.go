package notification

import (
	"context"

	"Momentum/application/telemetry"
)

func (service *Service) Send(ctx context.Context, input Input) (Response, error) {
	ctx, span := telemetry.Start(ctx, "application.notification.send")
	defer span.End()

	if service == nil || service.sender == nil {
		return Response{}, ErrUnavailable
	}

	if err := service.sender.Send(ctx, input); err != nil {
		return Response{}, err
	}

	return Response{Title: input.Title, Body: input.Body, Topic: input.Topic}, nil
}

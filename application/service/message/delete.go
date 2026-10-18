package message

import (
	"context"
	"errors"

	"Momentum/application/telemetry"
	domainmessage "Momentum/domain/message"
)

func (service *Service) Delete(ctx context.Context, id string) (MessageResponse, error) {
	ctx, span := telemetry.Start(ctx, "application.message.delete")
	defer span.End()

	if id == "" {
		return MessageResponse{}, ErrInvalidInput
	}

	if err := service.repository.Delete(ctx, id); err != nil {
		if errors.Is(err, domainmessage.ErrMessageNotFound) {
			return MessageResponse{}, ErrMessageNotFound
		}
		return MessageResponse{}, err
	}

	service.bumpMessageCacheVersion(ctx)
	return MessageResponse{}, nil
}

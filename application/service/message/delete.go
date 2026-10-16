package message

import (
	"context"
	"errors"

	domainmessage "Momentum/domain/message"
)

func (service *Service) Delete(ctx context.Context, id string) (MessageResponse, error) {
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

package message

import (
	"context"
	"errors"

	"Momentum/application/telemetry"
	domainmessage "Momentum/domain/message"
)

func (service *Service) Update(ctx context.Context, input UpdateInput) (MessageResponse, error) {
	ctx, span := telemetry.Start(ctx, "application.message.update")
	defer span.End()

	if input.ID == "" {
		return MessageResponse{}, ErrInvalidInput
	}

	passages := make([]domainmessage.UpdatePassageParams, 0, len(input.Passages))
	for _, passage := range input.Passages {
		if passage.ID == "" {
			return MessageResponse{}, ErrInvalidInput
		}
		passages = append(passages, domainmessage.UpdatePassageParams{
			ID:      passage.ID,
			Header:  passage.Header,
			Verse:   passage.Verse,
			Message: passage.Message,
		})
	}

	if err := service.repository.Update(ctx, domainmessage.UpdateParams{
		ID:        input.ID,
		Thumbnail: input.Thumbnail,
		Series:    input.Series,
		Title:     input.Title,
		Preacher:  input.Preacher,
		Date:      input.Date,
		Published: input.Published,
		Passages:  passages,
	}); err != nil {
		switch {
		case errors.Is(err, domainmessage.ErrMessageNotFound):
			return MessageResponse{}, ErrMessageNotFound
		case errors.Is(err, domainmessage.ErrPassageNotFound):
			return MessageResponse{}, ErrPassageNotFound
		default:
			return MessageResponse{}, err
		}
	}

	service.bumpMessageCacheVersion(ctx)
	return messageResponseFromUpdateInput(input), nil
}

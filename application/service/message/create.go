package message

import (
	"context"

	"Momentum/application/telemetry"
	domainmessage "Momentum/domain/message"
	"github.com/lucsky/cuid"
)

func (service *Service) Create(ctx context.Context, input CreateInput) (MessageResponse, error) {
	ctx, span := telemetry.Start(ctx, "application.message.create")
	defer span.End()

	passages := make([]domainmessage.CreatePassageParams, 0, len(input.Passages))
	for _, passage := range input.Passages {
		switch passage.Type {
		case headerPassageType:
			header := passage.Header
			passages = append(passages, domainmessage.CreatePassageParams{
				ID:     cuid.New(),
				Header: &header,
				Order:  passage.Order,
			})
		case messagePassageType:
			verse := passage.Verse
			messageText := passage.Message
			passages = append(passages, domainmessage.CreatePassageParams{
				ID:      cuid.New(),
				Verse:   &verse,
				Message: &messageText,
				Order:   passage.Order,
			})
		}
	}

	messageID := cuid.New()
	if err := service.repository.Create(ctx, domainmessage.CreateParams{
		ID:        messageID,
		Thumbnail: input.Thumbnail,
		Series:    input.Series,
		Title:     input.Title,
		Preacher:  input.Preacher,
		Date:      input.Date,
		Passages:  passages,
	}); err != nil {
		return MessageResponse{}, err
	}

	service.bumpMessageCacheVersion(ctx)

	return messageResponseFromCreateInput(input, messageID), nil
}

package message

import (
	"context"

	domainmessage "Momentum/domain/message"
	"github.com/lucsky/cuid"
)

func (service *Service) Create(ctx context.Context, input MessageInput) (MessageResponse, error) {
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

	if err := service.repository.Create(ctx, domainmessage.CreateParams{
		ID:        cuid.New(),
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
	return input, nil
}

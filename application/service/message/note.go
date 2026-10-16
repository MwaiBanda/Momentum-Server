package message

import (
	"context"
	"errors"

	domainmessage "Momentum/domain/message"
	"github.com/lucsky/cuid"
)

func (service *Service) CreateNote(ctx context.Context, input CreateNoteInput) (CreateNoteInput, error) {
	if input.PassageID == "" || input.UserID == "" {
		return CreateNoteInput{}, ErrInvalidInput
	}

	if err := service.repository.CreateNote(ctx, domainmessage.CreateNoteParams{
		ID:        cuid.New(),
		Content:   input.Content,
		UserID:    input.UserID,
		PassageID: input.PassageID,
	}); err != nil {
		return CreateNoteInput{}, err
	}

	service.bumpMessageCacheVersion(ctx)
	return input, nil
}

func (service *Service) UpdateNote(ctx context.Context, input NoteResponse) (NoteResponse, error) {
	if input.ID == "" {
		return NoteResponse{}, ErrInvalidInput
	}

	_, err := service.repository.UpdateNote(ctx, domainmessage.UpdateNoteParams{
		ID:      input.ID,
		Content: input.Content,
	})
	if err != nil {
		if errors.Is(err, domainmessage.ErrNoteNotFound) {
			return NoteResponse{}, ErrNoteNotFound
		}
		return NoteResponse{}, err
	}

	service.bumpMessageCacheVersion(ctx)
	return input, nil
}

func (service *Service) DeleteNote(ctx context.Context, id string) (NoteResponse, error) {
	if id == "" {
		return NoteResponse{}, ErrInvalidInput
	}

	deleted, err := service.repository.DeleteNote(ctx, id)
	if err != nil {
		if errors.Is(err, domainmessage.ErrNoteNotFound) {
			return NoteResponse{}, ErrNoteNotFound
		}
		return NoteResponse{}, err
	}

	service.bumpMessageCacheVersion(ctx)
	return toNoteResponse(deleted), nil
}

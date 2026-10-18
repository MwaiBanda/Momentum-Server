package message

import (
	"context"
	"errors"

	domainmessage "Momentum/domain/message"
	"github.com/lucsky/cuid"
)

func (service *Service) CreateNote(ctx context.Context, input CreateNoteInput) (CreateNoteResponse, error) {
	if input.PassageID == "" || input.UserID == "" {
		return CreateNoteResponse{}, ErrInvalidInput
	}

	id := cuid.New()
	if err := service.repository.CreateNote(ctx, domainmessage.CreateNoteParams{
		ID:        id,
		Content:   input.Content,
		UserID:    input.UserID,
		PassageID: input.PassageID,
	}); err != nil {
		return CreateNoteResponse{}, err
	}

	service.bumpMessageCacheVersion(ctx)
	return CreateNoteResponse{ID: id, Content: input.Content, UserID: input.UserID}, nil
}

func (service *Service) UpdateNote(ctx context.Context, input UpdateNoteInput) (NoteResponse, error) {
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
	return NoteResponse{ID: input.ID, Content: input.Content}, nil
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

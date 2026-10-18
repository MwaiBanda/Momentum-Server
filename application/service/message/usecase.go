package message

import "context"

type UseCase interface {
	ListPublishedByUser(ctx context.Context, userID string) (MessageListResponse, error)
	ListUnpublished(ctx context.Context) (MessageListResponse, error)
	Create(ctx context.Context, input CreateInput) (MessageResponse, error)
	Update(ctx context.Context, input UpdateInput) (MessageResponse, error)
	Delete(ctx context.Context, id string) (MessageResponse, error)
	CreateNote(ctx context.Context, input CreateNoteInput) (CreateNoteInput, error)
	UpdateNote(ctx context.Context, input NoteResponse) (NoteResponse, error)
	DeleteNote(ctx context.Context, id string) (NoteResponse, error)
}

package message

import "context"

type Repository interface {
	ListPublishedByUser(ctx context.Context, userID string) ([]Message, error)
	ListUnpublished(ctx context.Context) ([]Message, error)
	Create(ctx context.Context, params CreateParams) error
	Update(ctx context.Context, params UpdateParams) error
	Delete(ctx context.Context, id string) error
	CreateNote(ctx context.Context, params CreateNoteParams) error
	UpdateNote(ctx context.Context, params UpdateNoteParams) (Note, error)
	DeleteNote(ctx context.Context, id string) (Note, error)
}

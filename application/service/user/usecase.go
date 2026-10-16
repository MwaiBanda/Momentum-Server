package user

import "context"

type UseCase interface {
	GetByID(ctx context.Context, id string) (Response, error)
	Create(ctx context.Context, input CreateInput) (Response, error)
	Update(ctx context.Context, input UpdateInput) (Response, error)
	Delete(ctx context.Context, id string) (Response, error)
}

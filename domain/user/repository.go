package user

import "context"

type CreateParams struct {
	ID       string
	Email    string
	Fullname string
	Phone    string
}

type UpdateParams struct {
	ID       string
	Email    string
	Fullname string
	Phone    string
}

type Repository interface {
	GetByID(ctx context.Context, id string) (User, error)
	Create(ctx context.Context, params CreateParams) (User, error)
	Update(ctx context.Context, params UpdateParams) (User, error)
	Delete(ctx context.Context, id string) (User, error)
}

package transaction

import "context"

type Mailer interface {
	Send(ctx context.Context, input MailInput) error
}

type UseCase interface {
	List(ctx context.Context) ([]Response, error)
	ListByUser(ctx context.Context, userID string) ([]Response, error)
	Create(ctx context.Context, input Input) (Response, error)
	Delete(ctx context.Context, id string) (Response, error)
}

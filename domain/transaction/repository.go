package transaction

import "context"

type Repository interface {
	List(ctx context.Context) ([]Transaction, error)
	ListByUser(ctx context.Context, userID string) ([]Transaction, error)
	Create(ctx context.Context, params CreateParams) (Transaction, error)
	Delete(ctx context.Context, id string) (Transaction, error)
}

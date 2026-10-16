package sermon

import "context"

type Provider interface {
	Fetch(ctx context.Context, page string) ([]byte, error)
}

type UseCase interface {
	List(ctx context.Context, page string) (Response, error)
}

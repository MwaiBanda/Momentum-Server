package event

import "context"

type Provider interface {
	Fetch(ctx context.Context) ([]byte, error)
}

type UseCase interface {
	List(ctx context.Context) (Response, error)
}

package payment

import "context"

type Provider interface {
	Create(ctx context.Context, input Input) (ProviderResponse, error)
}

type ProviderResponse struct {
	Customer      string
	EphemeralKey  string
	PaymentIntent string
}

type UseCase interface {
	Create(ctx context.Context, input Input) (Response, error)
}

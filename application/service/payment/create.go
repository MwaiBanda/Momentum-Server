package payment

import (
	"context"

	"Momentum/application/telemetry"
)

func (service *Service) Create(ctx context.Context, input Input) (Response, error) {
	ctx, span := telemetry.Start(ctx, "application.payment.create")
	defer span.End()

	if service == nil || service.provider == nil || input.Amount <= 0 || input.Email == "" {
		return Response{}, ErrInvalidInput
	}

	created, err := service.provider.Create(ctx, input)
	if err != nil {
		return Response{}, err
	}

	return Response{
		PublishableKey: service.publishableKey,
		PaymentIntent:  created.PaymentIntent,
		Customer:       created.Customer,
		EphemeralKey:   created.EphemeralKey,
	}, nil
}

package payment

import (
	"context"
	"errors"
	"testing"
)

type fakeProvider struct {
	response ProviderResponse
	err      error
	input    Input
}

func (provider *fakeProvider) Create(_ context.Context, input Input) (ProviderResponse, error) {
	provider.input = input
	return provider.response, provider.err
}

func TestCreateMapsProviderResponse(t *testing.T) {
	provider := &fakeProvider{response: ProviderResponse{
		Customer:      "cus_123",
		EphemeralKey:  "ephkey",
		PaymentIntent: "pi_secret",
	}}
	service := GetInstance(provider, "pk_test")

	response, err := service.Create(context.Background(), Input{
		Amount: 100,
		Email:  "mwai@example.com",
	})
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}
	if response.PublishableKey != "pk_test" || response.PaymentIntent != "pi_secret" {
		t.Fatalf("response = %#v, want mapped payment response", response)
	}
}

func TestCreateRejectsInvalidInput(t *testing.T) {
	service := GetInstance(&fakeProvider{}, "pk_test")

	_, err := service.Create(context.Background(), Input{Amount: 0})
	if !errors.Is(err, ErrInvalidInput) {
		t.Fatalf("Create() error = %v, want %v", err, ErrInvalidInput)
	}
}

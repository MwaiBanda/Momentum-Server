package stripe

import (
	"context"
	"fmt"

	applicationpayment "Momentum/application/service/payment"

	stripeSDK "github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/customer"
	"github.com/stripe/stripe-go/v74/ephemeralkey"
	"github.com/stripe/stripe-go/v74/paymentintent"
)

type Client struct {
	version string
}

func NewClient(secretKey string, version string) *Client {
	stripeSDK.Key = secretKey
	return &Client{version: version}
}

func (client *Client) Create(ctx context.Context, input applicationpayment.Input) (applicationpayment.ProviderResponse, error) {
	if client == nil || stripeSDK.Key == "" {
		return applicationpayment.ProviderResponse{}, fmt.Errorf("Stripe client is not configured")
	}
	if err := ctx.Err(); err != nil {
		return applicationpayment.ProviderResponse{}, err
	}

	createdCustomer, err := customer.New(&stripeSDK.CustomerParams{
		Email: stripeSDK.String(input.Email),
		Name:  stripeSDK.String(input.Fullname),
		Phone: stripeSDK.String(input.Phone),
	})
	if err != nil {
		return applicationpayment.ProviderResponse{}, fmt.Errorf("create Stripe customer: %w", err)
	}

	keyParams := &stripeSDK.EphemeralKeyParams{
		Customer:      stripeSDK.String(createdCustomer.ID),
		StripeVersion: stripeSDK.String(client.version),
	}
	ephemeralKey, err := ephemeralkey.New(keyParams)
	if err != nil {
		return applicationpayment.ProviderResponse{}, fmt.Errorf("create Stripe ephemeral key: %w", err)
	}

	intentParams := &stripeSDK.PaymentIntentParams{
		Amount:       stripeSDK.Int64(int64(input.Amount)),
		Currency:     stripeSDK.String(string(stripeSDK.CurrencyUSD)),
		Customer:     stripeSDK.String(createdCustomer.ID),
		ReceiptEmail: stripeSDK.String(input.Email),
		AutomaticPaymentMethods: &stripeSDK.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripeSDK.Bool(true),
		},
	}
	paymentIntent, err := paymentintent.New(intentParams)
	if err != nil {
		return applicationpayment.ProviderResponse{}, fmt.Errorf("create Stripe payment intent: %w", err)
	}

	return applicationpayment.ProviderResponse{
		Customer:      createdCustomer.ID,
		EphemeralKey:  ephemeralKey.Secret,
		PaymentIntent: paymentIntent.ClientSecret,
	}, nil
}

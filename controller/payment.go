package controller

import (
	_ "Momentum/httputil"
	"Momentum/model"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v74"
	"log"
	"net/url"
	"os"
)

// PostPayment godoc
//
//	@Summary		Post a payment
//	@Description	Used to initiate stripe transactions
//	@tags			Payments
//	@Param			data	body	model.TransactionRequest	true	"Post a transaction"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.TransactionResponse
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/api/v1/payment [post]
func (controller *Controller) PostPayment(context *fiber.Ctx) error {
	transaction := new(model.TransactionRequest)
	if err := context.BodyParser(transaction); err != nil {
		log.Panic(err.Error())
	}

	params := url.Values{}
	params.Add("email", transaction.Email)
	params.Add("name", transaction.Fullname)
	params.Add("phone", transaction.Phone)

	customerResponse, err := controller.StripeRequest("/customers?" + params.Encode())
	if err != nil {
		log.Panic(err.Error())
	}
	customer := new(stripe.Customer)
	if err := json.Unmarshal(customerResponse, customer); err != nil {
		log.Panic(err)
	}

	params = url.Values{}
	params.Add("customer", customer.ID)

	ephemeralKeyResponse, err := controller.StripeRequest("/ephemeral_keys?" + params.Encode())
	if err != nil {
		log.Panic(err)
	}
	ephemeralKey := new(stripe.EphemeralKey)
	if err := json.Unmarshal(ephemeralKeyResponse, ephemeralKey); err != nil {
		log.Panic(err)
	}

	params = url.Values{}
	params.Add("customer", customer.ID)
	params.Add("amount", transaction.Amount)
	params.Add("receipt_email", transaction.Email)
	params.Add("currency", "usd")
	params.Add("automatic_payment_methods[enabled]", "true")

	paymentIntentResponse, err := controller.StripeRequest("/payment_intents?" + params.Encode())
	if err != nil {
		log.Panic(err)
	}
	paymentIntent := new(stripe.PaymentIntent)
	if err := json.Unmarshal(paymentIntentResponse, paymentIntent); err != nil {
		log.Panic(err)
	}

	return context.JSON(model.TransactionResponse{
		PublishableKey: os.Getenv("STRIPE_PUBLISHABLE_KEY"),
		Customer:       customer.ID,
		EphemeralKey:   ephemeralKey.Secret,
		PaymentIntent:  paymentIntent.ClientSecret,
	})
}

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
	"strconv"
	"sync"
)

// PostPayment godoc
//
//	@Summary		Post a payment
//	@Description	Used to initiate stripe transactions
//	@tags			Payments
//	@Param			Authorization	header	string					true	"Provide a bearer token"	example(Bearer XXX-xxx-XXX-xxx-XX)
//	@Param			data			body	model.PaymentRequest	true	"Post a transaction"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.PaymentResponse
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/api/v1/payments [post]
func (controller *Controller) PostPayment(context *fiber.Ctx) error {
	customer := new(stripe.Customer)
	paymentRequest := new(model.PaymentRequest)
	customerChannel := make(chan StripeResult)
	intentAndKeyChannel := make(chan StripeResult)
	waitGroup := new(sync.WaitGroup)

	if err := context.BodyParser(paymentRequest); err != nil {
		log.Panic(err.Error())
	}

	params := url.Values{}
	params.Add("email", paymentRequest.Email)
	params.Add("name", paymentRequest.Fullname)
	params.Add("phone", paymentRequest.Phone)
	waitGroup.Add(1)
	go controller.StripeRequest(waitGroup, customerChannel, "/customers?"+params.Encode())

	go func() {
		waitGroup.Wait()
		close(customerChannel)
	}()

	for result := range customerChannel {
		switch result.Endpoint {
		case "/customers":
			customerResponse, err := result.ReadData()
			if err != nil {
				log.Println(err.Error())
			}
			if err := json.Unmarshal(customerResponse, customer); err != nil {
				log.Panic(err)
			}

			params = url.Values{}
			params.Add("customer", customer.ID)
			waitGroup.Add(1)
			go controller.StripeRequest(waitGroup, intentAndKeyChannel, "/ephemeral_keys?"+params.Encode())

			params = url.Values{}
			params.Add("customer", customer.ID)
			params.Add("amount", strconv.Itoa(paymentRequest.Amount))
			params.Add("receipt_email", paymentRequest.Email)
			params.Add("currency", "usd")
			params.Add("automatic_payment_methods[enabled]", "true")
			waitGroup.Add(1)
			go controller.StripeRequest(waitGroup, intentAndKeyChannel, "/payment_intents?"+params.Encode())
		}
	}

	go func() {
		waitGroup.Wait()
		close(intentAndKeyChannel)
	}()

	ephemeralKey := new(stripe.EphemeralKey)
	paymentIntent := new(stripe.PaymentIntent)

	for result := range intentAndKeyChannel {
		switch result.Endpoint {
		case "/ephemeral_keys":
			ephemeralKeyResponse, err := result.ReadData()
			if err != nil {
				log.Panic(err)
			}
			if err := json.Unmarshal(ephemeralKeyResponse, ephemeralKey); err != nil {
				log.Panic(err)
			}
		case "/payment_intents":
			paymentIntentResponse, err := result.ReadData()
			if err != nil {
				log.Panic(err)
			}
			if err := json.Unmarshal(paymentIntentResponse, paymentIntent); err != nil {
				log.Panic(err)
			}
		}
	}

	return context.JSON(model.PaymentResponse{
		PublishableKey: os.Getenv("STRIPE_PUBLISHABLE_KEY"),
		Customer:       customer.ID,
		EphemeralKey:   ephemeralKey.Secret,
		PaymentIntent:  paymentIntent.ClientSecret,
	})
}

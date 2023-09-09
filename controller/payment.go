package controller

import (
	"Momentum/constants"
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
//	@Failure		400	{object}	model.HTTPError
//	@Failure		404	{object}	model.HTTPError
//	@Failure		500	{object}	model.HTTPError
//	@Router			/api/v1/payments [post]
func (controller *Controller) PostPayment(context *fiber.Ctx) error {
	paymentRequest := new(model.PaymentRequest)
	customer := new(stripe.Customer)

	customerChannel := make(chan StripeResponse)
	intentAndKeyChannel := make(chan StripeResponse)

	waitGroup := new(sync.WaitGroup)

	if err := context.BodyParser(paymentRequest); err != nil {
		log.Panic(err.Error())
	}

	params := url.Values{}
	params.Add("email", paymentRequest.Email)
	params.Add("name", paymentRequest.Fullname)
	params.Add("phone", paymentRequest.Phone)
	waitGroup.Add(1)
	go controller.StripeRequest(waitGroup, customerChannel, constants.StripeCustomerRoute+params.Encode())

	go func() {
		waitGroup.Wait()
		close(customerChannel)
	}()

	for response := range customerChannel {
		switch response.Endpoint {
		case constants.StripeCustomerRoute:
			customerResponse, err := response.ReadData()
			if err != nil {
				log.Println(err.Error())
			}
			if err := json.Unmarshal(customerResponse, customer); err != nil {
				log.Panic(err)
			}

			params = url.Values{}
			params.Add("customer", customer.ID)
			waitGroup.Add(1)
			go controller.StripeRequest(waitGroup, intentAndKeyChannel, constants.StripeEphemeralKeyRoute+params.Encode())

			params = url.Values{}
			params.Add("customer", customer.ID)
			params.Add("amount", strconv.Itoa(paymentRequest.Amount))
			params.Add("receipt_email", paymentRequest.Email)
			params.Add("currency", "usd")
			params.Add("automatic_payment_methods[enabled]", "true")
			waitGroup.Add(1)
			go controller.StripeRequest(waitGroup, intentAndKeyChannel, constants.StripePaymentIntentRoute+params.Encode())
		}
	}

	go func() {
		waitGroup.Wait()
		close(intentAndKeyChannel)
	}()

	ephemeralKey := new(stripe.EphemeralKey)
	paymentIntent := new(stripe.PaymentIntent)

	for response := range intentAndKeyChannel {
		switch response.Endpoint {
		case constants.StripeEphemeralKeyRoute:
			ephemeralKeyResponse, err := response.ReadData()
			if err != nil {
				log.Panic(err)
			}
			if err := json.Unmarshal(ephemeralKeyResponse, ephemeralKey); err != nil {
				log.Panic(err)
			}
		case constants.StripePaymentIntentRoute:
			paymentIntentResponse, err := response.ReadData()
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

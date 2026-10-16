package payment

import (
	"errors"
	"log"

	applicationpayment "Momentum/application/service/payment"

	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	payments applicationpayment.UseCase
}

func GetInstance(payments applicationpayment.UseCase) *Handler {
	return &Handler{payments: payments}
}

// Create godoc
//
//	@Summary		Create a payment
//	@Description	Create a Stripe payment intent for a contribution
//	@Accept			json
//	@Produce		json
//	@tags			Payments
//	@Param			Authorization	header	string	true	"Provide a bearer token"	example(******
//	@Param			data			body	payment.Input	true	"Create a payment"
//	@Success		200	{object}	payment.Response
//	@Failure		400	{object} http.ErrorResponse
//	@Failure		500	{object} http.ErrorResponse
//	@Router			/api/v1/payments [post]
func (handler *Handler) Create(ctx fiber.Ctx) error {
	if handler.payments == nil {
		log.Printf("[CreatePayment] handler dependencies are not initialized")
		return fiber.ErrInternalServerError
	}

	var input applicationpayment.Input
	if err := ctx.Bind().Body(&input); err != nil {
		return fiber.ErrBadRequest
	}

	response, err := handler.payments.Create(ctx.Context(), input)
	if err != nil {
		if errors.Is(err, applicationpayment.ErrInvalidInput) {
			return fiber.ErrBadRequest
		}
		log.Printf("[CreatePayment] %v", err)
		return fiber.ErrInternalServerError
	}
	return ctx.JSON(response)
}

package transaction

import (
	"log"

	applicationtransaction "Momentum/application/service/transaction"

	"github.com/gofiber/fiber/v3"
)

// Create godoc
//
//	@Summary		Post a transaction
//	@Description	Used to post a transaction
//	@Accept			json
//	@Produce		json
//	@tags			Transactions
//	@Param			Authorization	header	string	true	"Provide a bearer token"	example(******
//	@Security	BearerAuth
//	@Failure	401	{object} http.ErrorResponse
//	@Param			data			body	transaction.Input	true	"Post a transaction"
//	@Success		200	{object}	transaction.Response
//	@Failure		400	{object} http.ErrorResponse
//	@Failure		500	{object} http.ErrorResponse
//	@Router			/api/v1/transactions [post]
func (handler *Handler) Create(ctx fiber.Ctx) error {
	if handler.transactions == nil {
		log.Printf("[CreateTransaction] handler dependencies are not initialized")
		return fiber.ErrInternalServerError
	}

	var input applicationtransaction.Input
	if err := ctx.Bind().Body(&input); err != nil {
		return fiber.ErrBadRequest
	}

	response, err := handler.transactions.Create(ctx.Context(), input)
	if err != nil {
		return handler.mapError("CreateTransaction", err)
	}
	return ctx.JSON(response)
}

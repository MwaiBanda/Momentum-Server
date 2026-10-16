package transaction

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

// List godoc
//
//	@Summary		Show a list of all transactions
//	@Description	Get all transactions
//	@Accept			json
//	@Produce		json
//	@tags			Transactions
//	@Param			Authorization	header	string	true	"Provide a bearer token"	example(******
//	@Success		200	{array}	transaction.Response
//	@Failure		500	{object} http.ErrorResponse
//	@Router			/api/v1/transactions [get]
func (handler *Handler) List(ctx fiber.Ctx) error {
	if handler.transactions == nil {
		log.Printf("[ListTransactions] handler dependencies are not initialized")
		return fiber.ErrInternalServerError
	}

	response, err := handler.transactions.List(ctx.Context())
	if err != nil {
		return handler.mapError("ListTransactions", err)
	}
	return ctx.JSON(response)
}

// ListByUser godoc
//
//	@Summary		Show a user's transactions
//	@Description	Get transactions belonging to a user
//	@Accept			json
//	@Produce		json
//	@tags			Transactions
//	@Param			Authorization	header	string	true	"Provide a bearer token"	example(******
//	@Param			userId			path	string	true	"provide user Id"
//	@Success		200	{array}	transaction.Response
//	@Failure		400	{object} http.ErrorResponse
//	@Failure		500	{object} http.ErrorResponse
//	@Router			/api/v1/transactions/{userId} [get]
func (handler *Handler) ListByUser(ctx fiber.Ctx) error {
	if handler.transactions == nil {
		log.Printf("[ListTransactionsByUser] handler dependencies are not initialized")
		return fiber.ErrInternalServerError
	}

	response, err := handler.transactions.ListByUser(ctx.Context(), ctx.Params("userId"))
	if err != nil {
		return handler.mapError("ListTransactionsByUser", err)
	}
	return ctx.JSON(response)
}

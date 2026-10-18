package transaction

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

// Delete godoc
//
//	@Summary		Delete a transaction
//	@Description	Delete a transaction by ID
//	@Accept			json
//	@Produce		json
//	@tags			Transactions
//	@Param			Authorization	header	string	true	"Provide a bearer token"	example(******
//	@Security	BearerAuth
//	@Failure	401	{object} http.ErrorResponse
//	@Param			transactionId	path	string	true	"provide transaction Id"
//	@Success		200	{object}	transaction.Response
//	@Failure		400	{object} http.ErrorResponse
//	@Failure		404	{object} http.ErrorResponse
//	@Failure		500	{object} http.ErrorResponse
//	@Router			/api/v1/transactions/{transactionId} [delete]
func (handler *Handler) Delete(ctx fiber.Ctx) error {
	if handler.transactions == nil {
		log.Printf("[DeleteTransaction] handler dependencies are not initialized")
		return fiber.ErrInternalServerError
	}

	response, err := handler.transactions.Delete(ctx.Context(), ctx.Params("transactionId"))
	if err != nil {
		return handler.mapError("DeleteTransaction", err)
	}
	return ctx.JSON(response)
}

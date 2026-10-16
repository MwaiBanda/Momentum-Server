package message

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

// Delete godoc
//
//	@Summary		Delete message information
//	@Description	Delete a message and its passages and notes
//	@Accept			json
//	@Produce		json
//	@tags			Messages
//	@Param			Authorization	header	string	true	"Provide a bearer token"	example(******
//	@Param			messageId		path	string	true	"provide message Id"
//	@Success		200	{object}	message.MessageResponse
//	@Failure		400	{object}	model.HTTPError
//	@Failure		404	{object}	model.HTTPError
//	@Failure		500	{object}	model.HTTPError
//	@Router			/api/v1/messages/{messageId} [delete]
func (handler *Handler) Delete(ctx fiber.Ctx) error {
	if handler.messages == nil {
		log.Printf("[DeleteMessage] handler dependencies are not initialized")
		return fiber.ErrInternalServerError
	}

	response, err := handler.messages.Delete(ctx.Context(), ctx.Params("messageId"))
	if err != nil {
		return handler.mapError("DeleteMessage", err)
	}

	return ctx.JSON(response)
}

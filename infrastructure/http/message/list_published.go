package message

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

// ListPublishedByUser godoc
//
//	@Summary		Show published messages available for a user
//	@Description	get published messages and that user's passage notes
//	@Accept			json
//	@Produce		json
//	@tags			Messages
//	@Param			Authorization	header	string	true	"Provide a bearer token"	example(******
//	@Param			userId			path	string	true	"provide user Id"
//	@Success		200	{object}	message.MessageListResponse
//	@Failure		400	{object}	model.HTTPError
//	@Failure		500	{object}	model.HTTPError
//	@Router			/api/v1/messages/{userId} [get]
func (handler *Handler) ListPublishedByUser(ctx fiber.Ctx) error {
	if handler.messages == nil {
		log.Printf("[ListPublishedMessages] handler dependencies are not initialized")
		return fiber.ErrInternalServerError
	}

	response, err := handler.messages.ListPublishedByUser(ctx.Context(), ctx.Params("userId"))
	if err != nil {
		return handler.mapError("ListPublishedMessages", err)
	}

	return ctx.JSON(response)
}

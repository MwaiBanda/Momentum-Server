package message

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

// ListUnpublished godoc
//
//	@Summary		Show unpublished messages
//	@Description	get unpublished messages and passages
//	@Accept			json
//	@Produce		json
//	@tags			Messages
//	@Param			Authorization	header	string	true	"Provide a bearer token"	example(******
//	@Security	BearerAuth
//	@Failure	401	{object} http.ErrorResponse
//	@Success		200	{object}	message.MessageListResponse
//	@Failure		500	{object} http.ErrorResponse
//	@Router			/api/v1/messages/unpublished [get]
func (handler *Handler) ListUnpublished(ctx fiber.Ctx) error {
	if handler.messages == nil {
		log.Printf("[ListUnpublishedMessages] handler dependencies are not initialized")
		return fiber.ErrInternalServerError
	}

	response, err := handler.messages.ListUnpublished(ctx.Context())
	if err != nil {
		return handler.mapError("ListUnpublishedMessages", err)
	}

	return ctx.JSON(response)
}

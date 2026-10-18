package message

import (
	"log"

	applicationmessage "Momentum/application/service/message"

	"github.com/gofiber/fiber/v3"
)

// Create godoc
//
//	@Summary		Post a Message
//	@Description	Used to post a message
//	@tags			Messages
//	@Param			Authorization	header	string	true	"Provide a bearer token"	example(******
//	@Security	BearerAuth
//	@Failure	401	{object} http.ErrorResponse
//	@Param			data			body	message.CreateInput	true	"Post a message"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	message.MessageResponse
//	@Failure		400	{object} http.ErrorResponse
//	@Failure		500	{object} http.ErrorResponse
//	@Router			/api/v1/messages [post]
func (handler *Handler) Create(ctx fiber.Ctx) error {
	if handler.messages == nil {
		log.Printf("[CreateMessage] handler dependencies are not initialized")
		return fiber.ErrInternalServerError
	}

	var input applicationmessage.CreateInput
	if err := ctx.Bind().Body(&input); err != nil {
		return fiber.ErrBadRequest
	}

	response, err := handler.messages.Create(ctx.Context(), input)
	if err != nil {
		return handler.mapError("CreateMessage", err)
	}

	return ctx.JSON(response)
}

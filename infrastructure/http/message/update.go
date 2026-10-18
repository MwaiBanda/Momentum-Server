package message

import (
	"log"

	applicationmessage "Momentum/application/service/message"

	"github.com/gofiber/fiber/v3"
)

// Update godoc
//
//	@Summary		Update a Message
//	@Description	Used to update a message
//	@tags			Messages
//	@Param			Authorization	header	string	true	"Provide a bearer token"	example(******
//	@Param			data			body	message.UpdateInput	true	"Update a message"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	message.MessageResponse
//	@Failure		400	{object} http.ErrorResponse
//	@Failure		404	{object} http.ErrorResponse
//	@Failure		500	{object} http.ErrorResponse
//	@Router			/api/v1/messages [put]
func (handler *Handler) Update(ctx fiber.Ctx) error {
	if handler.messages == nil {
		log.Printf("[UpdateMessage] handler dependencies are not initialized")
		return fiber.ErrInternalServerError
	}

	var input applicationmessage.UpdateInput
	if err := ctx.Bind().Body(&input); err != nil {
		return fiber.ErrBadRequest
	}

	response, err := handler.messages.Update(ctx.Context(), input)
	if err != nil {
		return handler.mapError("UpdateMessage", err)
	}

	return ctx.JSON(response)
}

package message

import (
	"log"

	applicationmessage "Momentum/application/service/message"

	"github.com/gofiber/fiber/v3"
)

// UpdateNote godoc
//
//	@Summary		Update a Note in a Passage
//	@Description	Used to update a user note in a passage
//	@tags			Messages
//	@Param			Authorization	header	string	true	"Provide a bearer token"	example(******
//	@Param			data			body	message.NoteResponse	true	"Update a note"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	message.NoteResponse
//	@Failure		400	{object}	model.HTTPError
//	@Failure		404	{object}	model.HTTPError
//	@Failure		500	{object}	model.HTTPError
//	@Router			/api/v1/messages/notes [put]
func (handler *Handler) UpdateNote(ctx fiber.Ctx) error {
	if handler.messages == nil {
		log.Printf("[UpdateMessageNote] handler dependencies are not initialized")
		return fiber.ErrInternalServerError
	}

	var input applicationmessage.NoteResponse
	if err := ctx.Bind().Body(&input); err != nil {
		return fiber.ErrBadRequest
	}

	response, err := handler.messages.UpdateNote(ctx.Context(), input)
	if err != nil {
		return handler.mapError("UpdateMessageNote", err)
	}

	return ctx.JSON(response)
}

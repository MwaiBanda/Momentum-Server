package message

import (
	"log"

	applicationmessage "Momentum/application/service/message"
	"Momentum/infrastructure/http/request"

	"github.com/gofiber/fiber/v3"
)

// UpdateNote godoc
//
//	@Summary		Update a Note in a Passage
//	@Description	Used to update a user note in a passage
//	@tags			Messages
//	@Param			Authorization	header	string	true	"Provide a bearer token"	example(******
//	@Security	BearerAuth
//	@Failure	401	{object} http.ErrorResponse
//	@Param			data			body	message.UpdateNoteInput	true	"Update a note"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	message.NoteResponse
//	@Failure		400	{object} http.ErrorResponse
//	@Failure		404	{object} http.ErrorResponse
//	@Failure		500	{object} http.ErrorResponse
//	@Router			/api/v1/messages/notes [put]
func (handler *Handler) UpdateNote(ctx fiber.Ctx) error {
	if handler.messages == nil {
		log.Printf("[UpdateMessageNote] handler dependencies are not initialized")
		return fiber.ErrInternalServerError
	}

	input, err := request.BindAndValidate[applicationmessage.UpdateNoteInput](ctx)
	if err != nil {
		return err
	}

	response, err := handler.messages.UpdateNote(ctx.Context(), input)
	if err != nil {
		return handler.mapError("UpdateMessageNote", err)
	}

	return ctx.JSON(response)
}

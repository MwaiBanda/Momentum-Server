package message

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

// DeleteNote godoc
//
//	@Summary		Delete note information by Id
//	@Description	Delete a note's information by providing an Id
//	@Accept			json
//	@Produce		json
//	@tags			Messages
//	@Param			Authorization	header	string	true	"Provide a bearer token"	example(******
//	@Param			noteId			path	string	true	"provide note Id"
//	@Success		200	{object}	message.NoteResponse
//	@Failure		400	{object}	model.HTTPError
//	@Failure		404	{object}	model.HTTPError
//	@Failure		500	{object}	model.HTTPError
//	@Router			/api/v1/messages/notes/{noteId} [delete]
func (handler *Handler) DeleteNote(ctx fiber.Ctx) error {
	if handler.messages == nil {
		log.Printf("[DeleteMessageNote] handler dependencies are not initialized")
		return fiber.ErrInternalServerError
	}

	response, err := handler.messages.DeleteNote(ctx.Context(), ctx.Params("noteId"))
	if err != nil {
		return handler.mapError("DeleteMessageNote", err)
	}

	return ctx.JSON(response)
}

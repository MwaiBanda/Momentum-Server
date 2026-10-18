package message

import (
	"log"

	applicationmessage "Momentum/application/service/message"
	"Momentum/infrastructure/http/request"

	"github.com/gofiber/fiber/v3"
)

// CreateNote godoc
//
//	@Summary		Add a Note to a Passage
//	@Description	Used to add user notes to a passage
//	@tags			Messages
//	@Param			Authorization	header	string	true	"Provide a bearer token"	example(******
//	@Security	BearerAuth
//	@Failure	401	{object} http.ErrorResponse
//	@Param			data			body	message.CreateNoteInput	true	"Post a note"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	message.CreateNoteResponse
//	@Failure		400	{object} http.ErrorResponse
//	@Failure		500	{object} http.ErrorResponse
//	@Router			/api/v1/messages/notes [post]
func (handler *Handler) CreateNote(ctx fiber.Ctx) error {
	if handler.messages == nil {
		log.Printf("[CreateMessageNote] handler dependencies are not initialized")
		return fiber.ErrInternalServerError
	}

	input, err := request.BindAndValidate[applicationmessage.CreateNoteInput](ctx)
	if err != nil {
		return err
	}

	response, err := handler.messages.CreateNote(ctx.Context(), input)
	if err != nil {
		return handler.mapError("CreateMessageNote", err)
	}

	return ctx.JSON(response)
}

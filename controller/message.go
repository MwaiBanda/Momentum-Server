package controller

import (
	"Momentum/model"
	"Momentum/prisma/db"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// GetAllMessages godoc
//
//	@Summary		Show a list of messages available
//	@Description	get a list of messages available
//	@Accept			json
//	@Produce		json
//	@tags			Transactions
//	@Param			Authorization	header		string	true	"Provide a bearer token"	example(Bearer XXX-xxx-XXX-xxx-XX)
//	@Success		200				{array}		model.MessageResponse
//	@Failure		400				{object}	model.HTTPError
//	@Failure		404				{object}	model.HTTPError
//	@Failure		500				{object}	model.HTTPError
//	@Router			/api/v1/messages/{userId} [get]
func (controller *Controller) GetAllMessages(context *fiber.Ctx) error {
	var messages []model.MessageResponse
	res, err := controller.prisma.Message.FindMany().With(
		db.Message.Passages.Fetch().With(
			db.Passage.Notes.Fetch(
				db.Note.UserID.Equals(context.Params("userId")),
			),
		),
	).Exec(controller.context)
	if err != nil {
		fmt.Println(err)
	}

	result, _ := json.MarshalIndent(res, "", "  ")
	if err := json.Unmarshal(result, &messages); err != nil {
		fmt.Println(err)
	}

	return context.JSON(messages)
}


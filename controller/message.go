package controller

import (
	"Momentum/constants"
	"Momentum/model"
	"Momentum/prisma/db"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

// GetAllMessages godoc
//
//	@Summary		Show a list of messages available
//	@Description	get a list of messages available
//	@Accept			json
//	@Produce		json
//	@tags			Messages
//	@Param			Authorization	header		string	true	"Provide a bearer token"	example(Bearer XXX-xxx-XXX-xxx-XX)
//	@Success		200				{array}		model.MessageResponse
//	@Failure		400				{object}	model.HTTPError
//	@Failure		404				{object}	model.HTTPError
//	@Failure		500				{object}	model.HTTPError
//	@Router			/api/v1/messages/{userId} [get]
func (controller *Controller) GetAllMessages(context *fiber.Ctx) error {
	var messages []model.Message
	cachedMessages, err := controller.Redis.Get(controller.Context, constants.MessageKey).Result()
	if err == redis.Nil {
		res, err := controller.PrismaClient.Message.FindMany().With(
			db.Message.Passages.Fetch().OrderBy(
				db.Passage.CreatedOn.Order(db.SortOrderDesc),
			).With(
				db.Passage.Notes.Fetch(
					db.Note.UserID.Equals(context.Params("userId")),
				),
			),
		).OrderBy(
			db.Message.CreatedOn.Order(db.SortOrderDesc),
		).Exec(controller.Context)
		if err != nil {
			fmt.Println(err)
		}
	
		result, _ := json.MarshalIndent(res, "", "  ")
		if err := json.Unmarshal(result, &messages); err != nil {
			fmt.Println(err)
		}
		controller.Redis.Set(controller.Context, constants.MessageKey, string(result), time.Hour*6)
	} else if err != nil { 
		fmt.Println(err.Error())
	} else {
		if err := json.Unmarshal([]byte(cachedMessages), &messages); err != nil {
			fmt.Println(err)
		}
	}

	return context.JSON(model.MessageResponse{Data: messages})
}

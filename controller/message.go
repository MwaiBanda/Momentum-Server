package controller

import (
	"Momentum/constants"
	"Momentum/model"
	"Momentum/prisma/db"
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/lucsky/cuid"
	"github.com/redis/go-redis/v9"
	"github.com/steebchen/prisma-client-go/runtime/transaction"
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
	var unpublished []model.Message
	userId := context.Params("userId")
	cachedMessages, err := controller.Redis.Get(controller.Context, constants.MessageKey).Result()
	if err == redis.Nil {
		res, err := controller.PrismaClient.Message.FindMany(
			db.Message.Published.Equals(true),
		).With(
			db.Message.Passages.Fetch().OrderBy(
				db.Passage.CreatedOn.Order(db.SortOrderAsc),
			).With(
				db.Passage.Notes.Fetch(
					db.Note.UserID.Equals(userId),
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
		
		if userId != constants.Admin && messages != nil{
			for _, message := range messages {
				if message.HasOrder {
					sort.Slice(message.Passages, func(i, j int) bool {
						return message.Passages[i].Order < message.Passages[j].Order
					})
				}
			}
			controller.Redis.Set(controller.Context, constants.MessageKey, string(result), time.Hour*24)
		} else {
			res, err := controller.PrismaClient.Message.FindMany(
				db.Message.Published.Equals(false),
			).OrderBy(
				db.Message.CreatedOn.Order(db.SortOrderDesc),
			).Exec(controller.Context)
			if err != nil {
				fmt.Println(err)
			}
	
			result, _ := json.MarshalIndent(res, "", "  ")
			if err := json.Unmarshal(result, &unpublished); err != nil {
				fmt.Println(err)
			}
			messages = append(messages, unpublished...)
		}
	} else if err != nil {
		fmt.Println(err.Error())
	} else {
		if err := json.Unmarshal([]byte(cachedMessages), &messages); err != nil {
			fmt.Println(err)
		}
	}
	return context.JSON(model.MessageResponse{Data: messages})
}


// PostMessage godoc
//
//	@Summary		Post a Message
//	@Description	Used to post a message
//	@tags			Messages
//	@Param			Authorization	header	string			true	"Provide a bearer token"	example(Bearer XXX-xxx-XXX-xxx-XX)
//	@Param			data			body	model.Message	true	"Post a message"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.Message
//	@Failure		400	{object}	model.HTTPError
//	@Failure		404	{object}	model.HTTPError
//	@Failure		500	{object}	model.HTTPError
//	@Router			/api/v1/messages [post]
func (controller *Controller) PostMessage(context *fiber.Ctx) error {
	var message model.Message
	if err := context.BodyParser(&message); err != nil {
		return err
	}
	messageId := cuid.New()
	var transactions []transaction.Param

	go controller.Redis.Expire(controller.Context, constants.MessageKey, time.Second*0)

	transactions = append(transactions, controller.PrismaClient.Message.CreateOne(
		db.Message.ID.Set(messageId),
		db.Message.Thumbnail.Set(message.Thumbnail),
		db.Message.Series.Set(message.Series),
		db.Message.Title.Set(message.Title),
		db.Message.Preacher.Set(message.Preacher),
		db.Message.Date.Set(message.Date),
		db.Message.HasOrder.Set(true),
	).Tx())
	for _, passage := range message.Passages {
		if passage.Type == constants.Header {
			transactions = append(transactions, controller.PrismaClient.Passage.CreateOne(
				db.Passage.Header.Set(passage.Header),
				db.Passage.MessageID.Set(messageId),
				db.Passage.Order.Set(passage.Order),
			).Tx())
		} else if passage.Type == constants.Message {
			transactions = append(transactions, controller.PrismaClient.Passage.CreateOne(
				db.Passage.Verse.Set(passage.Verse),
				db.Passage.Message.Set(passage.Message),
				db.Passage.MessageID.Set(messageId),
				db.Passage.Order.Set(passage.Order),
			).Tx())
		}
	}
	err := controller.PrismaClient.Prisma.Transaction(transactions...).Exec(controller.Context)
	if err != nil {
		log.Panic(err.Error())
	}
	return context.JSON(message)
}

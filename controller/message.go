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
//	@Summary		Show a list of messages available for a user
//	@Description	get a list of messages available for a user
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
	userId := context.Params("userId")
	cachedMessages, err := controller.Redis.Get(controller.Context, constants.MessageKey+"-"+userId).Result()
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
		if messages != nil && userId != constants.Admin {
			controller.Redis.Set(controller.Context, constants.MessageKey+"-"+userId, string(result), time.Hour*24)
			controller.AddUserCacheSession(constants.MessageKey + "-" + userId)
		} 
	} else if err != nil {
		fmt.Println(err.Error())
	} else {
		if err := json.Unmarshal([]byte(cachedMessages), &messages); err != nil {
			fmt.Println(err)
		}
	}
	for _, message := range messages {
		if message.HasOrder {
			sort.Slice(message.Passages, func(i, j int) bool {
				return message.Passages[i].Order < message.Passages[j].Order
			})
		}
	}
	return context.JSON(model.MessageResponse{Data: messages})
}

// GetUnpublishedMessages godoc
//
//	@Summary		Show a list of unpublished messages available
//	@Description	get a list of unpublished messages available
//	@Accept			json
//	@Produce		json
//	@tags			Messages
//	@Param			Authorization	header		string	true	"Provide a bearer token"	example(Bearer XXX-xxx-XXX-xxx-XX)
//	@Success		200				{array}		model.MessageResponse
//	@Failure		400				{object}	model.HTTPError
//	@Failure		404				{object}	model.HTTPError
//	@Failure		500				{object}	model.HTTPError
//	@Router			/api/v1/messages/unpublished [get]
func (controller *Controller) GetUnpublishedMessages(context *fiber.Ctx) error {
	var unpublished []model.Message
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
	return context.JSON(unpublished)
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

// UpdateMessage godoc
//
//	@Summary		Update a Message
//	@Description	Used to update a message
//	@tags			Messages
//	@Param			Authorization	header	string			true	"Provide a bearer token"	example(Bearer XXX-xxx-XXX-xxx-XX)
//	@Param			data			body	model.Message	true	"Post a message"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.Message
//	@Failure		400	{object}	model.HTTPError
//	@Failure		404	{object}	model.HTTPError
//	@Failure		500	{object}	model.HTTPError
//	@Router			/api/v1/messages [put]
func (controller *Controller) UpdateMessage(context *fiber.Ctx) error {
	var message model.Message
	controller.DeleteUserCachedSessions()
	if err := context.BodyParser(&message); err != nil {
		return err
	}
	var transactions []transaction.Param
	transactions = append(transactions, controller.PrismaClient.Message.FindUnique(
		db.Message.ID.Equals(message.ID),
	).Update(
		db.Message.Title.Set(message.Title),
		db.Message.Preacher.Set(message.Preacher),
		db.Message.Date.Set(message.Date),
		db.Message.Thumbnail.Set(message.Thumbnail),
		db.Message.Series.Set(message.Series),
		db.Message.Published.Set(message.Published),
	).Tx())

	if len(message.Passages) > 0 {
		for _, passage := range message.Passages {
			transactions = append(transactions, controller.PrismaClient.Passage.FindUnique(
				db.Passage.ID.Equals(passage.ID),
			).Update(
				db.Passage.Header.Set(passage.Header),
				db.Passage.Verse.Set(passage.Verse),
				db.Passage.Message.Set(passage.Message),
			).Tx())
		}
	}

	err := controller.PrismaClient.Prisma.Transaction(transactions...).Exec(controller.Context)
	if err != nil {
		log.Panic(err.Error())
	}

	return context.JSON(message)
}

// AddNoteToPassage godoc
//
//	@Summary		Add a Note to a Passage
//	@Description	Used to add user notes to a passage
//	@tags			Messages
//	@Param			Authorization	header	string	true	"Provide a bearer token"	example(Bearer XXX-xxx-XXX-xxx-XX)
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.NoteRequest
//	@Failure		400	{object}	model.HTTPError
//	@Failure		404	{object}	model.HTTPError
//	@Failure		500	{object}	model.HTTPError
//	@Router			/api/v1/messages/notes [post]
func (controller *Controller) AddUserNoteToMessage(context *fiber.Ctx) error {
	note := new(model.NoteRequest)
	if err := context.BodyParser(note); err != nil {
		return err
	}

	_, err := controller.Redis.Expire(controller.Context, constants.MessageKey+"-"+note.UserID, time.Second*0).Result()
	if err != nil {
		return err
	}

	controller.PrismaClient.Note.CreateOne(
		db.Note.Content.Set(note.Content),
		db.Note.UserID.Set(note.UserID),
		db.Note.PassageID.Set(note.PassageID),
	).Exec(controller.Context)
	return context.JSON(note)
}

// UpdateNote godoc
//
//	@Summary		Update a Note in a Passage
//	@Description	Used to update a user notes in a passage
//	@tags			Messages
//	@Param			Authorization	header	string			true	"Provide a bearer token"	example(Bearer XXX-xxx-XXX-xxx-XX)
//	@Param			data			body	model.Message	true	"Post a note"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.Note
//	@Failure		400	{object}	model.HTTPError
//	@Failure		404	{object}	model.HTTPError
//	@Failure		500	{object}	model.HTTPError
//	@Router			/api/v1/messages/notes [put]
func (controller *Controller) UpdateUserNote(context *fiber.Ctx) error {
	note := new(model.Note)
	if err := context.BodyParser(note); err != nil {
		return err
	}

	_, err := controller.Redis.Expire(controller.Context, constants.MessageKey+"-"+note.UserID, time.Second*0).Result()
	if err != nil {
		return err
	}
	controller.PrismaClient.Note.FindUnique(
		db.Note.ID.Equals(note.ID),
	).Update(
		db.Note.Content.Set(note.Content),
	).Exec(controller.Context)
	return context.JSON(note)
}


// DeleteNote godoc
//
//	@Summary		Delete note information by Id
//	@Description	Delete a note's information by providing an Id
//	@Accept			json
//	@Produce		json
//	@tags			Messages
//	@Param			Authorization	header		string	true	"Provide a bearer token"	example(Bearer XXX-xxx-XXX-xxx-XX)
//	@Param			noteId			path		string	true	"provide note Id"
//	@Success		200				{array}		model.Note
//	@Failure		400				{object}	model.HTTPError
//	@Failure		404				{object}	model.HTTPError
//	@Failure		500				{object}	model.HTTPError
//	@Router			/api/v1/messages/notes/{noteId} [delete]
func (controller *Controller) DeleteNote(context *fiber.Ctx) error {
	var note model.Note
	res, err := controller.PrismaClient.Note.FindUnique(
		db.Note.ID.Equals(context.Params("noteId")),
	).Delete().Exec(controller.Context)
	
	if err != nil {
		fmt.Println(err)
	}

	result, _ := json.MarshalIndent(res, "", "  ")
	if err := json.Unmarshal(result, &note); err != nil {
		fmt.Println(err)
	}

	if _, err := controller.Redis.Expire(controller.Context, constants.MessageKey+"-"+note.UserID, time.Second*0).Result(); err != nil {
		fmt.Println(err)
	}
	
	return context.JSON(note)
}

// DeleteMessage godoc
//
//	@Summary		Delete message information 
//	@Description	Delete a message's information 
//	@Accept			json
//	@Produce		json
//	@tags			Messages
//	@Param			Authorization	header		string			true	"Provide a bearer token"	example(Bearer XXX-xxx-XXX-xxx-XX)
//	@Param			data			body		model.Message	true	"Post a message"
//	@Success		200				{object}	model.Message
//	@Failure		400				{object}	model.HTTPError
//	@Failure		404				{object}	model.HTTPError
//	@Failure		500				{object}	model.HTTPError
//	@Router			/api/v1/messages [delete]
func (controller *Controller) DeleteMessage(context *fiber.Ctx) error {
	var message model.Message
	var transactions []transaction.Param
	for _, passage := range message.Passages {
		transactions = append(transactions, controller.PrismaClient.Passage.FindUnique(
			db.Passage.ID.Equals(passage.ID),
		).Delete().Tx())
	}
	transactions = append(transactions,  controller.PrismaClient.Message.FindUnique(
		db.Message.ID.Equals(context.Params("messageId")),
	).Delete().Tx())

	err := controller.PrismaClient.Prisma.Transaction(transactions...).Exec(controller.Context)
	if err != nil {
		log.Println(err.Error())
	}
	return context.JSON(message)
}

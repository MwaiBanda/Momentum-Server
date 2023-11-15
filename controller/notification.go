package controller

import (
	"Momentum/model"
	"log"

	"firebase.google.com/go/v4/messaging"
	"github.com/gofiber/fiber/v2"
)

// PostNotification godoc
//
//	@Summary		Post a notification
//	@Description	Used by clients to request a notification to be sent to a topiv
//	@tags			Notifications
//	@Param			Authorization	header	string				true	"Provide a bearer token"	example(Bearer XXX-xxx-XXX-xxx-XX)
//	@Param			data			body	model.Notification	true	"Post a notification"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.Notification
//	@Failure		400	{object}	model.HTTPError
//	@Failure		404	{object}	model.HTTPError
//	@Failure		500	{object}	model.HTTPError
//	@Router			/api/v1/notifications [post]
func (controller *Controller) PostNotification(context *fiber.Ctx) error {
	notification := new(model.Notification)

	if err := context.BodyParser(notification); err != nil {
		log.Panic(err.Error())
	}

	var sendNotification = func() error {
		message := &messaging.Message{
			Notification: &messaging.Notification{
				Title: notification.Title,
				Body:  notification.Body,
			},
			Topic: notification.Topic,
		}
		_, err := controller.Messaging.Send(controller.Context, message)
		if err != nil {
			log.Println(err.Error())
		}
		return context.JSON(notification)
	}

	if controller.Messaging != nil {
		return sendNotification()
	} else {
		for canSendNotification := range controller.CanSendNotificationChannel {
			if canSendNotification {
				close(controller.CanSendNotificationChannel)
				return sendNotification()
			}
		}
	}

	return context.SendStatus(500)
}

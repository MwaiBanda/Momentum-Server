package notification

import (
	"log"

	applicationnotification "Momentum/application/service/notification"

	"github.com/gofiber/fiber/v3"
)

// Send godoc
//
//	@Summary		Post a notification
//	@Description	Used by clients to request a notification to be sent to a topic
//	@tags			Notifications
//	@Param			Authorization	header	string	true	"Provide a bearer token"	example(******
//	@Param			data			body	notification.Input	true	"Post a notification"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	notification.Input
//	@Failure		400	{object} http.ErrorResponse
//	@Failure		500	{object} http.ErrorResponse
//	@Router			/api/v1/notifications [post]
func (handler *Handler) Send(ctx fiber.Ctx) error {
	if handler.notifications == nil {
		log.Printf("[SendNotification] handler dependencies are not initialized")
		return fiber.ErrInternalServerError
	}

	var input applicationnotification.Input
	if err := ctx.Bind().Body(&input); err != nil {
		return fiber.ErrBadRequest
	}

	response, err := handler.notifications.Send(ctx.Context(), input)
	if err != nil {
		return handler.mapError("SendNotification", err)
	}

	return ctx.JSON(response)
}

package event

import (
	"log"

	applicationevent "Momentum/application/service/event"

	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	events applicationevent.UseCase
}

func GetInstance(events applicationevent.UseCase) *Handler {
	return &Handler{events: events}
}

// List godoc
//
//	@Summary		Get all calendar events
//	@Description	Get all events scheduled on the church calendar
//	@tags			Events
//	@Param			Authorization	header	string	true	"Provide a bearer token"	example(******
//	@Security	BearerAuth
//	@Failure	401	{object} http.ErrorResponse
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	event.Response
//	@Failure		500	{object} http.ErrorResponse
//	@Router			/api/v1/events [get]
func (handler *Handler) List(ctx fiber.Ctx) error {
	if handler.events == nil {
		log.Printf("[ListEvents] handler dependencies are not initialized")
		return fiber.ErrInternalServerError
	}

	response, err := handler.events.List(ctx.Context())
	if err != nil {
		log.Printf("[ListEvents] %v", err)
		return fiber.ErrInternalServerError
	}

	return ctx.JSON(response)
}

package message

import (
	"errors"
	"log"

	applicationmessage "Momentum/application/service/message"

	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	messages applicationmessage.UseCase
}

func GetInstance(messages applicationmessage.UseCase) *Handler {
	return &Handler{messages: messages}
}

func (handler *Handler) mapError(operation string, err error) error {
	switch {
	case errors.Is(err, applicationmessage.ErrInvalidInput):
		return fiber.ErrBadRequest
	case errors.Is(err, applicationmessage.ErrMessageNotFound),
		errors.Is(err, applicationmessage.ErrPassageNotFound),
		errors.Is(err, applicationmessage.ErrNoteNotFound):
		return fiber.ErrNotFound
	default:
		log.Printf("[%s] %v", operation, err)
		return fiber.ErrInternalServerError
	}
}

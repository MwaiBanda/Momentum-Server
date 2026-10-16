package notification

import (
	"errors"
	"log"

	applicationnotification "Momentum/application/service/notification"

	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	notifications applicationnotification.UseCase
}

func GetInstance(notifications applicationnotification.UseCase) *Handler {
	return &Handler{notifications: notifications}
}

func (handler *Handler) mapError(operation string, err error) error {
	if errors.Is(err, applicationnotification.ErrUnavailable) {
		log.Printf("[%s] %v", operation, err)
		return fiber.ErrInternalServerError
	}

	log.Printf("[%s] %v", operation, err)
	return fiber.ErrInternalServerError
}

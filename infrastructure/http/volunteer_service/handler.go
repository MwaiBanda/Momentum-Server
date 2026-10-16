package volunteer_service

import (
	"errors"
	"log"

	applicationservice "Momentum/application/service/volunteer_service"

	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	services applicationservice.UseCase
}

func GetInstance(services applicationservice.UseCase) *Handler {
	return &Handler{services: services}
}

func (handler *Handler) mapError(operation string, err error) error {
	switch {
	case errors.Is(err, applicationservice.ErrInvalidInput):
		return fiber.ErrBadRequest
	case errors.Is(err, applicationservice.ErrServiceNotFound),
		errors.Is(err, applicationservice.ErrDayNotFound):
		return fiber.ErrNotFound
	default:
		log.Printf("[%s] %v", operation, err)
		return fiber.ErrInternalServerError
	}
}

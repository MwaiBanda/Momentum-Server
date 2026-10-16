package meal

import (
	"errors"
	"log"

	applicationmeal "Momentum/application/service/meal"

	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	meals applicationmeal.UseCase
}

func GetInstance(meals applicationmeal.UseCase) *Handler {
	return &Handler{meals: meals}
}

func (handler *Handler) mapError(operation string, err error) error {
	switch {
	case errors.Is(err, applicationmeal.ErrInvalidInput):
		return fiber.ErrBadRequest
	case errors.Is(err, applicationmeal.ErrNotFound),
		errors.Is(err, applicationmeal.ErrVolunteeredMealNotFound):
		return fiber.ErrNotFound
	default:
		log.Printf("[%s] %v", operation, err)
		return fiber.ErrInternalServerError
	}
}

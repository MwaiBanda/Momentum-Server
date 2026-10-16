package transaction

import (
	"errors"
	"log"

	applicationtransaction "Momentum/application/service/transaction"

	"github.com/gofiber/fiber/v3"
)

type Handler struct {
	transactions applicationtransaction.UseCase
}

func GetInstance(transactions applicationtransaction.UseCase) *Handler {
	return &Handler{transactions: transactions}
}

func (handler *Handler) mapError(operation string, err error) error {
	switch {
	case errors.Is(err, applicationtransaction.ErrInvalidInput):
		return fiber.ErrBadRequest
	case errors.Is(err, applicationtransaction.ErrNotFound):
		return fiber.ErrNotFound
	default:
		log.Printf("[%s] %v", operation, err)
		return fiber.ErrInternalServerError
	}
}

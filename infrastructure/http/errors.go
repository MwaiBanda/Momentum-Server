package http

import (
	"errors"
	"log"
	stdhttp "net/http"

	"github.com/gofiber/fiber/v3"
)

type ErrorResponse struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"Bad Request"`
}

func ErrorHandler(ctx fiber.Ctx, err error) error {
	status := fiber.StatusInternalServerError
	message := stdhttp.StatusText(status)

	var fiberError *fiber.Error
	if errors.As(err, &fiberError) {
		status = fiberError.Code
		message = fiberError.Message
	} else {
		log.Printf("[HTTP] %v", err)
	}

	if message == "" {
		message = stdhttp.StatusText(status)
	}

	return ctx.Status(status).JSON(ErrorResponse{
		Code:    status,
		Message: message,
	})
}

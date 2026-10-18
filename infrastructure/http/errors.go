package http

import (
	"errors"
	"log"
	stdhttp "net/http"

	applicationerrors "Momentum/application/errors"
	"Momentum/infrastructure/http/request"

	"github.com/gofiber/fiber/v3"
)

type ErrorResponse struct {
	Code      int    `json:"code" example:"400" binding:"required"`
	Error     string `json:"error" example:"invalid_request" binding:"required"`
	Message   string `json:"message" example:"Bad Request" binding:"required"`
	RequestID string `json:"requestId,omitempty" example:"a1b2c3d4"`
}

func ErrorHandler(ctx fiber.Ctx, err error) error {
	status := fiber.StatusInternalServerError
	message := stdhttp.StatusText(status)
	errorCode := applicationerrors.CodeInternal
	knownError := false

	if applicationCode, applicationStatus, applicationMessage, ok := applicationerrors.Details(err); ok {
		errorCode = applicationCode
		status = applicationStatus
		message = applicationMessage
		knownError = true
	} else {
		var fiberError *fiber.Error
		if errors.As(err, &fiberError) {
			status = fiberError.Code
			message = fiberError.Message
			errorCode = errorCodeForStatus(status)
			knownError = true
		} else {
			log.Printf("[HTTP] request_id=%s error=%v", request.FromContext(ctx), err)
		}
	}

	if knownError && status >= fiber.StatusInternalServerError {
		log.Printf("[HTTP] request_id=%s error=%v", request.FromContext(ctx), err)
	}

	if message == "" {
		message = stdhttp.StatusText(status)
	}

	return ctx.Status(status).JSON(ErrorResponse{
		Code:      status,
		Error:     errorCode,
		Message:   message,
		RequestID: request.FromContext(ctx),
	})
}

func errorCodeForStatus(status int) string {
	switch status {
	case fiber.StatusBadRequest:
		return applicationerrors.CodeInvalidRequest
	case fiber.StatusUnauthorized:
		return applicationerrors.CodeUnauthorized
	case fiber.StatusForbidden:
		return applicationerrors.CodeForbidden
	case fiber.StatusNotFound:
		return applicationerrors.CodeNotFound
	case fiber.StatusConflict:
		return applicationerrors.CodeConflict
	default:
		if status >= fiber.StatusInternalServerError {
			return applicationerrors.CodeInternal
		}
		return applicationerrors.CodeInvalidRequest
	}
}

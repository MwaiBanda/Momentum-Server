package request

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v3"
)

const (
	RequestIDHeader = "X-Request-ID"
	requestIDKey    = "momentum.request_id"
)

func RequestID() fiber.Handler {
	return func(ctx fiber.Ctx) error {
		requestID := normalizeRequestID(ctx.Get(RequestIDHeader))
		if requestID == "" {
			var bytes [16]byte
			if _, err := rand.Read(bytes[:]); err != nil {
				return fmt.Errorf("generate request ID: %w", err)
			}
			requestID = hex.EncodeToString(bytes[:])
		}

		ctx.Locals(requestIDKey, requestID)
		ctx.Response().Header.Set(RequestIDHeader, requestID)
		return ctx.Next()
	}
}

func FromContext(ctx fiber.Ctx) string {
	requestID, _ := ctx.Locals(requestIDKey).(string)
	return requestID
}

func normalizeRequestID(requestID string) string {
	requestID = strings.TrimSpace(requestID)
	if requestID == "" || len(requestID) > 128 {
		return ""
	}

	for _, character := range requestID {
		if (character >= 'a' && character <= 'z') ||
			(character >= 'A' && character <= 'Z') ||
			(character >= '0' && character <= '9') ||
			character == '-' ||
			character == '_' ||
			character == '.' {
			continue
		}
		return ""
	}

	return requestID
}

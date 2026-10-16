package middleware

import (
	"crypto/sha256"
	"crypto/subtle"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/keyauth"
)

func BearerTokenAuthentication(apiKey string) fiber.Handler {
	return keyauth.New(keyauth.Config{
		Validator: func(c fiber.Ctx, key string) (bool, error) {
			hashedAPIKey := sha256.Sum256([]byte(apiKey))
			hashedKey := sha256.Sum256([]byte(key))

			if subtle.ConstantTimeCompare(hashedAPIKey[:], hashedKey[:]) == 1 {
				return true, nil
			}
			return false, keyauth.ErrMissingOrMalformedAPIKey
		},
	})
}

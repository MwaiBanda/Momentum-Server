package middleware

import (
	"crypto/sha256"
	"crypto/subtle"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/keyauth"
)

func BearerTokenAuthentication() fiber.Handler {
	return keyauth.New(keyauth.Config{
		Validator: func(c fiber.Ctx, key string) (bool, error) {
			hashedAPIKey := sha256.Sum256([]byte(os.Getenv("API_KEY")))
			hashedKey := sha256.Sum256([]byte(key))

			if subtle.ConstantTimeCompare(hashedAPIKey[:], hashedKey[:]) == 1 {
				return true, nil
			}
			return false, keyauth.ErrMissingOrMalformedAPIKey
		},
	})
}

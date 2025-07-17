package middlewares

import (
	"strings"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func NewJWTMiddleware(secret string) fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			Key: []byte(secret),
		},
		SuccessHandler: func(c *fiber.Ctx) error {
			authorizationHeader := c.Get("Authorization")

			if authorizationHeader == "" {
				return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized: Missing Authorization header")
			}

			authParts := strings.Fields(authorizationHeader)
			if len(authParts) != 2 || authParts[0] != "Bearer" {
				return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized: Invalid Authorization header format")
			}

			token := authParts[1]
			c.Locals("jwtToken", token)

			return c.Next()
		},
	})
}

package middlewares

import (
	"fmt"

	"integration-auth-service/configs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func NewBasicAuth(env *configs.Configs) fiber.Handler {
	return basicauth.New(basicauth.Config{
		Users: map[string]string{
			env.Auth.BasicAuthUsername: env.Auth.BasicAuthPassword,
		},
		Authorizer: func(user, pass string) bool {
			fmt.Println("Authenticating:", user)
			return user == env.Auth.BasicAuthUsername && pass == env.Auth.BasicAuthPassword
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
		},
	})
}

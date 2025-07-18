package middlewares

import (
	"integration-auth-service/pkg/loggers"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SystemLoggerMiddleware(logger loggers.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		logger.SystemLogger(c, start, err)
		return err
	}
}

func DbLoggerMiddleware(logger loggers.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		err := logger.DbLogger(c)
		if err != nil {
			return err
		}
		return nil
	}
}

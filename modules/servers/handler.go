package servers

import (
	_authControllers "integration-auth-service/modules/auth/controllers"
	_authRepositories "integration-auth-service/modules/auth/repositories"
	_authUsecases "integration-auth-service/modules/auth/usecases"
	"integration-auth-service/modules/middlewares"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func (s *Server) MapHandlers() error {

	s.App.Use(middlewares.RecoverMiddleware())

	// Swagger UI
	s.App.Get("/swagger/*", fiberSwagger.WrapHandler)

	s.App.Get("/health-check", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// Group a version
	v1 := s.App.Group("/v1")

	// Public routes
	publicGroup := v1.Group("/integration-api")

	// Auth Controller
	authRepository := _authRepositories.NewAuthRepository(s.Db, s.C)
	authUsecase := _authUsecases.NewAuthUsecase(s.Cfg, authRepository)
	_authControllers.NewAuthController(publicGroup, authUsecase)

	// End point not found response
	s.App.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     "error, end point not found",
			"result":      nil,
		})
	})

	return nil
}

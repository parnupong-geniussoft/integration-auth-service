package servers

import (
	_authControllers "integration-auth-service/modules/auth/controllers"
	_authRepositories "integration-auth-service/modules/auth/repositories"
	_authUsecases "integration-auth-service/modules/auth/usecases"
	"integration-auth-service/modules/middlewares"

	_ "integration-auth-service/docs"

	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func (s *Server) MapHandlers() error {
	// Swagger UI
	s.App.Get("/swagger/*", fiberSwagger.WrapHandler)

	s.App.Get("/health-check", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"status": "ok"})
	})

	s.App.Use(middlewares.SystemLoggerMiddleware(*s.Log))
	s.App.Use(middlewares.DbLoggerMiddleware(*s.Log))
	s.App.Use(middlewares.RecoverMiddleware())

	// Group a version
	v1 := s.App.Group("/v1")

	// Public routes
	publicGroup := v1.Group("/integration-api")

	// Auth Controller
	authRepository := _authRepositories.NewAuthRepository(s.Db, s.C)
	authUsecase := _authUsecases.NewAuthUsecase(s.Cfg, authRepository)
	_authControllers.NewAuthController(publicGroup, authUsecase)

	// End point not found response
	s.App.Use(func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     "error, end point not found",
			"result":      nil,
		})
	})

	return nil
}

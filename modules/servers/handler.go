package servers

import (
	_userControllers "go-fiber-clean-arch-example/modules/users/controllers"
	_userRepositories "go-fiber-clean-arch-example/modules/users/repositories"
	_userUsecases "go-fiber-clean-arch-example/modules/users/usecases"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) MapHandlers() error {
	// Group a version
	v1 := s.App.Group("/v1")

	//* Users group
	usersGroup := v1.Group("/users")
	usersRepository := _userRepositories.NewUsersRepository(s.Db)
	usersUsecase := _userUsecases.NewUsersUsecase(usersRepository)
	_userControllers.NewUsersController(usersGroup, usersUsecase)

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

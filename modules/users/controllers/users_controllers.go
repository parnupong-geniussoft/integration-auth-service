package controllers

import (
	"go-fiber-clean-arch-example/modules/entities"
	"go-fiber-clean-arch-example/modules/users/usecases"

	"github.com/gofiber/fiber/v2"
)

type usersController struct {
	UsersUsecase usecases.UsersUsecase
}

func NewUsersController(r fiber.Router, usersUsecase usecases.UsersUsecase) {
	controllers := &usersController{
		UsersUsecase: usersUsecase,
	}
	r.Post("/", controllers.Register)
}

func (h *usersController) Register(c *fiber.Ctx) error {
	req := new(entities.UsersRegisterReq)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{
			"status":      fiber.ErrBadRequest.Message,
			"status_code": fiber.ErrBadRequest.Code,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	res, err := h.UsersUsecase.Register(req)
	if err != nil {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     err.Error(),
			"result":      nil,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "OK",
		"status_code": fiber.StatusOK,
		"message":     "",
		"result":      res,
	})
}

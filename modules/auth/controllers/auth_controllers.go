package controllers

import (
	"integration-auth-service/modules/auth/usecases"
	"integration-auth-service/modules/entities"

	"github.com/gofiber/fiber/v2"
)

type authController struct {
	AuthUsecase usecases.AuthUsecase
}

func NewAuthController(r fiber.Router, authUsecase usecases.AuthUsecase) {
	controllers := &authController{
		AuthUsecase: authUsecase,
	}
	r.Post("/request_token", controllers.RequestToken)
}

func (h *authController) RequestToken(c *fiber.Ctx) error {
	req := new(entities.TokenRequest)

	if err := c.BodyParser(req); err != nil {
		return err
	}

	resp, err := h.AuthUsecase.GetToken(req)

	if err != nil {
		c.Context().SetStatusCode(fiber.StatusUnauthorized)
		return fiber.NewError(401, err.Error())
	}

	return c.JSON(resp)
}

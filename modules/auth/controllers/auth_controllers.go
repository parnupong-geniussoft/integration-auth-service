package controllers

import (
	"integration-auth-service/modules/auth/entities"
	"integration-auth-service/modules/auth/usecases"

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

// RequestToken godoc
// @Summary Request OAuth Token
// @Description รับ client_id และ client_secret เพื่อขอ JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body entities.TokenRequest true "Token request body"
// @Success 200 {object} entities.TokenResponse
// @Router /v1/integration-api/request_token [post]
func (h *authController) RequestToken(c *fiber.Ctx) error {
	req := new(entities.TokenRequest)

	if err := c.BodyParser(req); err != nil {
		return err
	}

	resp, err := h.AuthUsecase.GetToken(req)

	if err != nil {
		c.Context().SetStatusCode(fiber.StatusUnauthorized)
		return fiber.NewError(fiber.StatusUnauthorized, err.Error())
	}

	return c.JSON(resp)
}

package handlers

import (
	"github.com/asifhajiyev/matching-api/model"
	"github.com/asifhajiyev/matching-api/services"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type AuthHandler interface {
	GetToken(c *fiber.Ctx) error
}

type authHandler struct {
	authService services.AuthService
}

func NewAuthHandler(as services.AuthService) AuthHandler {
	return authHandler{authService: as}
}

func (ah authHandler) GetToken(c *fiber.Ctx) error {
	r, err := ah.authService.GetToken()

	if err != nil {
		return c.Status(err.Code).JSON(
			model.BuildRestResponse(err.Code, err.Message, nil, err.Details))
	}
	return c.Status(http.StatusOK).JSON(
		model.BuildRestResponse(http.StatusOK, http.StatusText(http.StatusOK), r, nil))
}

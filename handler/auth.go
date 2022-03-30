package handler

import (
	"github.com/asifhajiyev/matching-api/model/response"
	"github.com/asifhajiyev/matching-api/service"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type AuthHandler interface {
	GetToken(c *fiber.Ctx) error
}

type authHandler struct {
	As service.AuthService
}

func NewAuthHandler(as service.AuthService) AuthHandler {
	return authHandler{As: as}
}

func (ah authHandler) GetToken(c *fiber.Ctx) error {

	r, err := ah.As.GetToken()

	if err != nil {
		return c.Status(err.Code).JSON(response.RestResponse{
			Code:    err.Code,
			Message: err.Message,
			Data:    nil,
		})
	}
	return c.Status(http.StatusOK).JSON(response.RestResponse{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
		Data:    r,
	})
}

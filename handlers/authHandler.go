package handlers

import (
	"github.com/asifhajiyev/matching-api/logger"
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

// GetToken godoc
// @Summary			Get Token to call match/* endpoints
// @Tags 			Auth
// @Description 	Generating Bearer token
// @Produce     	json
// @Success     	200  {object}  model.RestResponse
// @Router      	/auth/get-token [get]
func (ah authHandler) GetToken(c *fiber.Ctx) error {
	logger.Info("GetToken.begin")
	r, err := ah.authService.GetToken()

	if err != nil {
		logger.Error("GetToken.error", err)
		return c.Status(err.Code).JSON(
			model.BuildRestResponse(err.Code, err.Message, nil, err.Details))
	}
	logger.Info("GetToken.end")
	return c.Status(http.StatusOK).JSON(
		model.BuildRestResponse(http.StatusOK, http.StatusText(http.StatusOK), r, nil))
}

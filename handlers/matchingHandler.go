package handlers

import (
	"github.com/asifhajiyev/matching-api/model/response"
	"github.com/asifhajiyev/matching-api/services"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type MatchingHandler interface {
	Match(c *fiber.Ctx) error
}

type matchingHandler struct {
	Ms services.MatchingService
}

func NewMatchingHandler(ms services.MatchingService) MatchingHandler {
	return matchingHandler{Ms: ms}
}

func (mh matchingHandler) Match(c *fiber.Ctx) error {
	longitude := c.Query("longitude")
	latitude := c.Query("latitude")

	r, err := mh.Ms.Match(longitude, latitude)

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

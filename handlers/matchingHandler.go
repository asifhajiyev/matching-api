package handlers

import (
	"github.com/asifhajiyev/matching-api/model"
	"github.com/asifhajiyev/matching-api/services"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type MatchingHandler interface {
	Match(c *fiber.Ctx) error
}

type matchingHandler struct {
	service services.MatchingService
}

func NewMatchingHandler(ms services.MatchingService) MatchingHandler {
	return matchingHandler{service: ms}
}

func (mh matchingHandler) Match(c *fiber.Ctx) error {
	longitude := c.Query("longitude")
	latitude := c.Query("latitude")

	r, err := mh.service.Match(longitude, latitude)

	if err != nil {
		return c.Status(err.Code).JSON(
			model.BuildRestResponse(err.Code, err.Message, nil, err.Details))
	}
	return c.Status(http.StatusOK).JSON(
		model.BuildRestResponse(http.StatusOK, http.StatusText(http.StatusOK), r, nil))
}

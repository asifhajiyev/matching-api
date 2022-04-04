package handlers

import (
	"github.com/asifhajiyev/matching-api/logger"
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

// Match godoc
// @Summary			Match Rider with the nearest driver
// @Tags 			Match
// @Description 	Matches given rider with the nearest driver by calculating distance
// @Accept      	json
// @Produce     	json
// @Success     	200  {object}  model.RestResponse
// @securityDefinitions.bearerAuth
// @Security 		bearerAuth
// @In 				header
// @Name 			Bearer
// @Param 			Authorization header string true "Bearer"
// @Param       	longitude 	query string true "longitude of rider"
// @Param       	latitude  	query string true "latitude of rider"
// @Router      	/match [get]
func (mh matchingHandler) Match(c *fiber.Ctx) error {
	logger.Info("Match.begin")
	longitude := c.Query("longitude")
	latitude := c.Query("latitude")

	r, err := mh.service.Match(longitude, latitude)

	if err != nil {
		logger.Error("Match.error", err)
		return c.Status(err.Code).JSON(
			model.BuildRestResponse(err.Code, err.Message, nil, err.Details))
	}
	logger.Info("Match.end", r)
	return c.Status(http.StatusOK).JSON(
		model.BuildRestResponse(http.StatusOK, http.StatusText(http.StatusOK), r, nil))
}

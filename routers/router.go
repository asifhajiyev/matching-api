package routers

import (
	"github.com/asifhajiyev/matching-api/handlers"
	"github.com/asifhajiyev/matching-api/middleware"
	"github.com/asifhajiyev/matching-api/model/response"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"net/http"
)

type HandlerList struct {
	Mh handlers.MatchingHandler
	Ah handlers.AuthHandler
}

func (h *HandlerList) SetupRoutes(app *fiber.App) {
	app.Use(logger.New())

	appBaseGroup := app.Group("api")

	mh := appBaseGroup.Group("match").Use(middleware.JWTProtector())
	h.SetupMatchingRoute(mh)

	ah := appBaseGroup.Group("auth")
	h.SetupAuthRoute(ah)

	handleNotFoundError(app)
}

func handleNotFoundError(app *fiber.App) {
	app.Use(
		func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusNotFound).JSON(response.RestResponse{
				Code:    fiber.StatusNotFound,
				Message: http.StatusText(fiber.StatusNotFound),
				Data:    nil,
			})
		},
	)
}

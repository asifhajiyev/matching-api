package routers

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/asifhajiyev/matching-api/constants"
	"github.com/asifhajiyev/matching-api/handlers"
	"github.com/asifhajiyev/matching-api/middleware"
	"github.com/asifhajiyev/matching-api/model"
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
	useSwagger(app)

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
			return c.Status(fiber.StatusNotFound).JSON(
				model.BuildRestResponse(fiber.StatusNotFound, http.StatusText(fiber.StatusNotFound),
					nil, constants.ErrorURLNotFound))
		},
	)
}

func useSwagger(app *fiber.App) {
	route := app.Group("/swagger")
	route.Get("*", swagger.HandlerDefault)
}

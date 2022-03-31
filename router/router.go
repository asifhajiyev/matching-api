package router

import (
	"github.com/asifhajiyev/matching-api/handler"
	"github.com/asifhajiyev/matching-api/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type HandlerList struct {
	Mh handler.MatchingHandler
	Ah handler.AuthHandler
}

func (h *HandlerList) SetupRoutes(app *fiber.App) {
	app.Use(logger.New())

	appBaseGroup := app.Group("api")

	mh := appBaseGroup.Group("match").Use(middleware.JWTProtected())
	h.SetupMatchingRoute(mh)

	ah := appBaseGroup.Group("auth")
	h.SetupAuthRoute(ah)
}

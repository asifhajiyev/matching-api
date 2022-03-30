package router

import (
	"github.com/asifhajiyev/matching-api/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type HandlerList struct {
	Mh handler.MatchingHandler
	Ah handler.AuthHandler
}

func (h *HandlerList) SetupRoutes(app *fiber.App) {
	app.Use(logger.New())

	mh := app.Group("api").Group("match")
	h.SetupMatchingRoute(mh)

	ah := app.Group("api").Group("auth")
	h.SetupAuthRoute(ah)
}

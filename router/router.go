package router

import (
	"github.com/asifhajiyev/matching-api/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type HandlerList struct {
	Mh handler.MatchingHandler
}

func (h *HandlerList) SetupRoutes(app *fiber.App) {
	app.Use(logger.New())

	dl := app.Group("api").Group("match")
	h.SetupMatchingRoute(dl)
}

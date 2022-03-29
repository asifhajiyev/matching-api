package router

import "github.com/gofiber/fiber/v2"

func (h HandlerList) SetupMatchingRoute(r fiber.Router) {
	r.Get("/match", h.Mh.Match)
}

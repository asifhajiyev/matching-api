package routers

import "github.com/gofiber/fiber/v2"

func (h HandlerList) SetupMatchingRoute(r fiber.Router) {
	r.Get("/", h.Mh.Match)
}

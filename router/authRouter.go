package router

import "github.com/gofiber/fiber/v2"

func (h HandlerList) SetupAuthRoute(r fiber.Router) {
	r.Get("/get-token", h.Ah.GetToken)
}

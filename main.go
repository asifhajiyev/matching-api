package main

import (
	"github.com/asifhajiyev/matching-api/clients"
	"github.com/asifhajiyev/matching-api/handler"
	"github.com/asifhajiyev/matching-api/router"
	"github.com/asifhajiyev/matching-api/service"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {

	InitEnvVariables()

	resty := resty.New()
	resty.SetBaseURL("http://localhost:8080/api/")
	matchClient := clients.NewDriverClient(resty)
	matchService := service.NewMatchingService(matchClient)
	matchHandler := handler.NewMatchingHandler(matchService)

	app := fiber.New()
	r := router.HandlerList{Mh: matchHandler, Ah: handler.NewAuthHandler(service.JwtAuthService{})}
	r.SetupRoutes(app)

	log.Fatal(app.Listen(":8090"))
}

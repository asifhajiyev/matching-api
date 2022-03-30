package main

import (
	"github.com/asifhajiyev/matching-api/client"
	"github.com/asifhajiyev/matching-api/handler"
	"github.com/asifhajiyev/matching-api/router"
	"github.com/asifhajiyev/matching-api/service"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {

	resty := resty.New()
	resty.SetBaseURL("http://localhost:8080/api/")
	matchClient := client.NewDriverClient(resty)
	matchService := service.NewMatchingService(matchClient)
	matchHandler := handler.NewMatchingHandler(matchService)

	app := fiber.New()
	r := router.HandlerList{Mh: matchHandler}
	r.SetupRoutes(app)

	log.Fatal(app.Listen(":8090"))
}

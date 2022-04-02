package main

import (
	"github.com/asifhajiyev/matching-api/clients"
	"github.com/asifhajiyev/matching-api/handlers"
	"github.com/asifhajiyev/matching-api/routers"
	"github.com/asifhajiyev/matching-api/services"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {

	InitEnvVariables()

	r := resty.New()
	r.SetBaseURL("http://localhost:8080/api/")

	matchClient := clients.NewDriverClient(r)
	matchService := services.NewMatchingService(matchClient)
	matchHandler := handlers.NewMatchingHandler(matchService)

	authHandler := handlers.NewAuthHandler(services.JwtAuthService{})

	app := fiber.New()
	hl := routers.HandlerList{
		Mh: matchHandler,
		Ah: authHandler,
	}
	hl.SetupRoutes(app)

	log.Fatal(app.Listen(":8090"))
}

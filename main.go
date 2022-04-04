package main

import (
	"github.com/asifhajiyev/matching-api/clients"
	_ "github.com/asifhajiyev/matching-api/docs"
	"github.com/asifhajiyev/matching-api/handlers"
	"github.com/asifhajiyev/matching-api/routers"
	"github.com/asifhajiyev/matching-api/services"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"log"
)

// @title           Driver Location API
// @version         1.0
// @description     This is a Driver Location API to save them and search

// @contact.email  	asif.hajiyev@outlook.com
// @BasePath  /api/

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

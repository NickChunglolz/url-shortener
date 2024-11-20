package main

import (
	"fmt"
	"log"

	"github.com/NickChunglolz/url-shortener-command/adapter"
	"github.com/NickChunglolz/url-shortener-command/infrastructure/repository"
	"github.com/NickChunglolz/url-shortener-command/main/utils"
	"github.com/NickChunglolz/url-shortener-command/usecase"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: false,
	}))

	var factory utils.DatabaseFactory = utils.NewDatabaseFactory()
	defer factory.CloseDatabaseConnections()

	db, err := factory.CreateDb()
	if err != nil {
		fmt.Println("DB connection error:", err)
	}

	shortenedUrlRepository := repository.NewShortenedUrlRepositoryImpl(db)
	counterRepository := repository.NewCounterRepositoryImpl(db)
	shortenedUrlCommand := usecase.NewShortenedUrlCommand(shortenedUrlRepository, counterRepository)
	controller := adapter.NewController(shortenedUrlCommand)
	controller.SetRoutes(app)

	log.Fatal(app.Listen("localhost:8082"))
}

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
	config, err := utils.NewConfig().Load()
	if err != nil {
		fmt.Println("Load system config error:", err)
	}

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{fmt.Sprintf("http://%s:%s", config.Upstream.Server.Host, config.Upstream.Server.Port)},
		AllowMethods:     []string{"POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: false,
	}))

	var factory utils.DatabaseFactory = utils.NewDatabaseFactory(config)
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

	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port)))
}

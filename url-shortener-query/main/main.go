package main

import (
	"fmt"
	"log"

	"github.com/NickChunglolz/url-shortener-query/adapter"
	"github.com/NickChunglolz/url-shortener-query/infrastructure/repository"
	"github.com/NickChunglolz/url-shortener-query/main/utils"
	"github.com/NickChunglolz/url-shortener-query/usecase"
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

	var factory utils.DatabaseFactoryInterface = utils.NewDatabaseFactory(config)
	defer factory.CloseDatabaseConnections()

	db, err := factory.CreateDb()
	if err != nil {
		fmt.Println("DB connection error:", err)
	}

	cacheDb, err := factory.CreateCacheDb()
	if err != nil {
		fmt.Println("CacheDB connection error:", err)
	}

	shortenedUrlRepository := repository.NewShortenedUrlRepositoryImpl(config, db, cacheDb)
	shortenedUrlCommand := usecase.NewShortenedUrlQuery(shortenedUrlRepository)
	controller := adapter.NewController(shortenedUrlCommand)
	controller.SetRoutes(app)

	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port)))
}

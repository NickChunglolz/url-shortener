package main

import (
	"fmt"
	"log"

	"github.com/NickChunglolz/url-shortener-query/adapter"
	"github.com/NickChunglolz/url-shortener-query/infrastructure/repository"
	"github.com/NickChunglolz/url-shortener-query/main/utils"
	"github.com/NickChunglolz/url-shortener-query/usecase"
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	var factory utils.DatabaseFactoryInterface = utils.NewDatabaseFactory()
	defer factory.CloseDatabaseConnections()

	db, err := factory.CreateDb()
	if err != nil {
		fmt.Println("DB connection error:", err)
	}

	cacheDb, err := factory.CreateCacheDb()
	if err != nil {
		fmt.Println("CacheDB connection error:", err)
	}

	shortenedUrlRepository := repository.NewShortenedUrlRepositoryImpl(db, cacheDb)
	shortenedUrlCommand := usecase.NewShortenedUrlQuery(shortenedUrlRepository)
	controller := adapter.NewController(shortenedUrlCommand)
	controller.SetRoutes(app)

	log.Fatal(app.Listen("localhost:8081"))
}

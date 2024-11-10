package main

import (
	"log"

	"github.com/NickChunglolz/url-shortener-query/main/utils"
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	var factory utils.DatabaseFactoryInterface = utils.NewDatabaseFactory()
	defer factory.CloseDatabaseConnections()

	log.Fatal(app.Listen("localhost:8081"))
}

package main

import (
	"log"

	"github.com/NickChunglolz/url-shortener-command/main/utils"
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	var factory utils.DatabaseFactory = utils.NewDatabaseFactory()
	defer factory.CloseDatabaseConnections()

	factory.CreateDb()

	log.Fatal(app.Listen("localhost:8082"))
}

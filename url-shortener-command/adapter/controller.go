package adapter

import (
	"github.com/NickChunglolz/url-shortener-command/usecase"
	"github.com/gofiber/fiber/v3"
)

const rootPath = "/url-shortener-command/Urls"

type Controller struct {
	command *usecase.ShortenedUrlCommand
}

func NewController(command *usecase.ShortenedUrlCommand) *Controller {
	return &Controller{
		command: command,
	}
}

func (controller *Controller) setRoutes(app *fiber.App) {
	app.Post(rootPath, func(c fiber.Ctx) error {
		var request usecase.CreateShortenUrlRequest
		if err := c.Bind().JSON(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		if request.OriginalURL == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Original URL is required",
			})
		}

		data, err := controller.command.CreateShortenUrl(&request)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(data)
	})
}

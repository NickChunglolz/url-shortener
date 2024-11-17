package adapter

import (
	"fmt"

	"github.com/NickChunglolz/url-shortener-query/usecase"
	"github.com/gofiber/fiber/v3"
)

const rootPath = "/Urls"

type Controller struct {
	query *usecase.ShortenedUrlQuery
}

func NewController(query *usecase.ShortenedUrlQuery) *Controller {
	return &Controller{
		query: query,
	}
}

func (controller *Controller) SetRoutes(app *fiber.App) {
	app.Get("/:code", func(c fiber.Ctx) error {
		code := c.Params("code")

		data, err := controller.query.GetShortenUrlByCode(code)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		if data == nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Short URL not found",
			})
		}
	
		return c.Status(fiber.StatusFound).Redirect().To(data.LongUrl)
	})

	app.Get(fmt.Sprintf("%s/:code", rootPath), func(c fiber.Ctx) error {
		code := c.Params("code")

		data, err := controller.query.GetShortenUrlByCode(code)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(data)
	})

	app.Get(rootPath, func(c fiber.Ctx) error {
		data, err := controller.query.QueryShortenUrls()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(fiber.StatusOK).JSON(data)
	})
}
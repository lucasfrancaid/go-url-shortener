package fiber_router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucasfrancaid/go-url-shortener/pkg/application/dto"
	"github.com/lucasfrancaid/go-url-shortener/pkg/port/controller"
)

func Stats(c *fiber.Ctx) error {
	shortenedURL := c.Params("shortenedURL")
	d := dto.ShortenedDTO{ShortenedURL: shortenedURL}

	ctl := controller.NewShortenerController()
	pre := ctl.Stats(d)
	res := pre.HTTP()

	return c.Status(res.StatusCode).JSON(res.Data)
}

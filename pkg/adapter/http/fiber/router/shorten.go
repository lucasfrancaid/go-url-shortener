package fiber_router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucasfrancaid/go-url-shortener/pkg/application/dto"
	"github.com/lucasfrancaid/go-url-shortener/pkg/port/controller"
)

func Shorten(c *fiber.Ctx) error {
	d := new(dto.ShortenDTO)
	if err := c.BodyParser(d); err != nil {
		return err
	}

	ctl := controller.NewShortenerController()
	pre := ctl.Shorten(*d)
	res := pre.HTTP()

	return c.Status(res.StatusCode).JSON(res.Data)
}

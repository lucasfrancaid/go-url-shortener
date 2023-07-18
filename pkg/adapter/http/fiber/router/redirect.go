package fiber_router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lucasfrancaid/go-url-shortener/pkg/application/dto"
	"github.com/lucasfrancaid/go-url-shortener/pkg/port/controller"
)

func Redirect(c *fiber.Ctx) error {
	hashedURL := c.Params("hashedURL")
	d := dto.ShortenedDTO{ShortenedURL: hashedURL}

	ctl := controller.NewShortenerController()
	pre := ctl.Redirect(d)
	res := pre.HTTP()

	if pre.Error == nil {
		if data, ok := res.Data.(dto.RedirectDTO); ok {
			return c.Redirect(data.URL, res.StatusCode)
		}
	}

	return c.Status(res.StatusCode).JSON(res.Data)
}

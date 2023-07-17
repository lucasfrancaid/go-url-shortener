package echo_router

import (
	"github.com/labstack/echo/v4"
	"github.com/lucasfrancaid/go-url-shortener/pkg/application/dto"
	"github.com/lucasfrancaid/go-url-shortener/pkg/port/controller"
)

func Stats(c echo.Context) error {
	shortenedURL := c.Param("shortenedURL")
	d := dto.ShortenedDTO{ShortenedURL: shortenedURL}

	ctl := controller.NewShortenerController()
	pre := ctl.Stats(d)
	res := pre.HTTP()

	return c.JSON(res.StatusCode, res.Data)
}

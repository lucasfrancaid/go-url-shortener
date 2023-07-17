package echo_router

import (
	"github.com/labstack/echo/v4"
	"github.com/lucasfrancaid/go-url-shortener/pkg/application/dto"
	"github.com/lucasfrancaid/go-url-shortener/pkg/port/controller"
)

func Redirect(c echo.Context) error {
	shortenedURL := c.Param("shortenedURL")
	d := dto.ShortenedDTO{ShortenedURL: shortenedURL}

	ctl := controller.NewShortenerController()
	pre := ctl.Redirect(d)
	res := pre.HTTP()

	if pre.Error == nil {
		if data, ok := res.Data.(dto.RedirectDTO); ok {
			return c.Redirect(res.StatusCode, data.URL)
		}
	}

	return c.JSON(res.StatusCode, res.Data)
}

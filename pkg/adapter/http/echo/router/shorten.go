package echo_router

import (
	"github.com/labstack/echo/v4"
	factory "github.com/lucasfrancaid/go-url-shortener/internal/pkg/infrastructure/factory/repository"
	"github.com/lucasfrancaid/go-url-shortener/pkg/application/dto"
	"github.com/lucasfrancaid/go-url-shortener/pkg/port/controller"
)

func Shorten(c echo.Context) error {
	d := new(dto.ShortenDTO)
	if err := (&echo.DefaultBinder{}).BindBody(c, &d); err != nil {
		return err
	}

	shortenerRepo := factory.NewShortenerRepository()
	statsRepo := factory.NewShortenerStatsRepository()
	ctl := controller.NewShortenerController(shortenerRepo, statsRepo)
	pre := ctl.Shorten(*d)
	res := pre.HTTP()

	return c.JSON(res.StatusCode, res.Data)
}

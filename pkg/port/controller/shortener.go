package controller

import (
	"github.com/lucasfrancaid/go-url-shortener/pkg/application/dto"
	"github.com/lucasfrancaid/go-url-shortener/pkg/application/usecase"
	"github.com/lucasfrancaid/go-url-shortener/pkg/port/presenter"
)

type ShortenerController struct{}

func NewShortenerController() ShortenerController {
	return ShortenerController{}
}

func (c *ShortenerController) Shorten(d dto.ShortenDTO) presenter.Presenter {
	u := usecase.NewShortenUseCase()
	r, err := u.Do(d)
	if err != nil {
		return presenter.PresenterError(err)
	}
	return presenter.PresenterSuccess(r)
}

func (c *ShortenerController) Redirect(d dto.ShortenedDTO) presenter.Presenter {
	u := usecase.NewRedirectUseCase()
	r, err := u.Do(d)
	if err != nil {
		return presenter.PresenterError(err)
	}
	return presenter.PresenterRedirect(r)
}

func (c *ShortenerController) Stats(d dto.ShortenedDTO) presenter.Presenter {
	u := usecase.NewStatsUseCase()
	r, err := u.Do(d)
	if err != nil {
		return presenter.PresenterError(err)
	}
	return presenter.PresenterSuccess(r)
}

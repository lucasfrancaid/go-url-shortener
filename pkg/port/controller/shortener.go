package controller

import (
	"github.com/lucasfrancaid/go-url-shortener/pkg/application/dto"
	"github.com/lucasfrancaid/go-url-shortener/pkg/application/usecase"
	"github.com/lucasfrancaid/go-url-shortener/pkg/port/presenter"
	"github.com/lucasfrancaid/go-url-shortener/pkg/port/repository"
)

type ShortenerController struct {
	shortenerRepository repository.ShortenerRepository
}

func NewShortenerController(shortenerRepository repository.ShortenerRepository) ShortenerController {
	return ShortenerController{shortenerRepository: shortenerRepository}
}

func (c *ShortenerController) Shorten(d dto.ShortenDTO) presenter.Presenter {
	u := usecase.NewShortenUseCase(c.shortenerRepository)
	r, err := u.Do(d)
	if err != nil {
		return presenter.PresenterError(err)
	}
	return presenter.PresenterSuccess(r)
}
package controller

import (
	"github.com/lucasfrancaid/go-url-shortener/pkg/application/dto"
	"github.com/lucasfrancaid/go-url-shortener/pkg/application/usecase"
	"github.com/lucasfrancaid/go-url-shortener/pkg/port/presenter"
	"github.com/lucasfrancaid/go-url-shortener/pkg/port/repository"
)

type ShortenerController struct {
	shortenerRepository repository.ShortenerRepository
	statsRepository     repository.ShortenerStatsRepository
}

func NewShortenerController(shortenerRepository repository.ShortenerRepository, statsRepository repository.ShortenerStatsRepository) ShortenerController {
	return ShortenerController{shortenerRepository: shortenerRepository, statsRepository: statsRepository}
}

func (c *ShortenerController) Shorten(d dto.ShortenDTO) presenter.Presenter {
	u := usecase.NewShortenUseCase(c.shortenerRepository, c.statsRepository)
	r, err := u.Do(d)
	if err != nil {
		return presenter.PresenterError(err)
	}
	return presenter.PresenterSuccess(r)
}

func (c *ShortenerController) Redirect(d dto.ShortenedDTO) presenter.Presenter {
	u := usecase.NewRedirectUseCase(c.shortenerRepository, c.statsRepository)
	r, err := u.Do(d)
	if err != nil {
		return presenter.PresenterError(err)
	}
	return presenter.PresenterRedirect(r)
}

func (c *ShortenerController) Stats(d dto.ShortenedDTO) presenter.Presenter {
	u := usecase.NewStatsUseCase(c.statsRepository)
	r, err := u.Do(d)
	if err != nil {
		return presenter.PresenterError(err)
	}
	return presenter.PresenterSuccess(r)
}

package usecase

import (
	"errors"
	"strings"

	factory "github.com/lucasfrancaid/go-url-shortener/internal/pkg/infrastructure/factory/repository"
	base "github.com/lucasfrancaid/go-url-shortener/pkg/application/base"
	"github.com/lucasfrancaid/go-url-shortener/pkg/application/dto"
	"github.com/lucasfrancaid/go-url-shortener/pkg/port/repository"
)

type RedirectUseCase struct {
	ShortenerRepository repository.ShortenerRepository
	StatsRepository     repository.ShortenerStatsRepository
}

func NewRedirectUseCase() RedirectUseCase {
	return RedirectUseCase{
		ShortenerRepository: factory.NewShortenerRepository(),
		StatsRepository:     factory.NewShortenerStatsRepository(),
	}
}

func (u *RedirectUseCase) Do(d dto.ShortenedDTO) (dto.RedirectDTO, error) {
	if err := u.validate(&d); err != nil {
		return dto.RedirectDTO{}, &base.Error{Type: base.VALIDATOR_ERROR, Err: err}
	}

	entity, err := u.ShortenerRepository.Read(d.ShortenedURL)
	if err != nil {
		if baseErr, ok := err.(*base.Error); ok {
			err = baseErr
		} else if strings.Contains(err.Error(), "not found") {
			err = &base.Error{Type: base.NOT_FOUND_ERROR, Err: err}
		} else {
			err = &base.Error{Type: base.INTERNAL_ERROR, Err: err}
		}
		return dto.RedirectDTO{}, err
	}

	u.StatsRepository.Increment(d.ShortenedURL)

	return dto.RedirectDTO{URL: entity.URL}, nil
}

func (u *RedirectUseCase) validate(d *dto.ShortenedDTO) error {
	if len(d.ShortenedURL) != 8 {
		return errors.New("HashedURL provided is invalid")
	}
	return nil
}

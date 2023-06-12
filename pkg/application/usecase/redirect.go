package usecase

import (
	"errors"

	base "github.com/lucasfrancaid/go-url-shortener/pkg/application/base"
	"github.com/lucasfrancaid/go-url-shortener/pkg/application/dto"
	"github.com/lucasfrancaid/go-url-shortener/pkg/port/repository"
)

type RedirectUseCase struct {
	shortenerRepository repository.ShortenerRepository
}

func NewRedirectUseCase(shortenerRepository repository.ShortenerRepository) RedirectUseCase {
	return RedirectUseCase{shortenerRepository: shortenerRepository}
}

func (u *RedirectUseCase) Do(d dto.ShortenedDTO) (dto.RedirectDTO, error) {
	if err := u.validate(&d); err != nil {
		return dto.RedirectDTO{}, &base.Error{Type: base.VALIDATOR_ERROR, Err: err}
	}
	entity, err := u.shortenerRepository.Read(d.ShortenedURL)
	if err != nil {
		return dto.RedirectDTO{}, &base.Error{Type: base.NOT_FOUND_ERROR, Err: err}
	}
	return dto.RedirectDTO{URL: entity.URL}, nil
}

func (u *RedirectUseCase) validate(d *dto.ShortenedDTO) error {
	if len(d.ShortenedURL) != 8 {
		return errors.New("HashedURL provided is invalid")
	}
	return nil
}

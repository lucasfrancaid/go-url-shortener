package usecase

import (
	"errors"
	"strings"

	base "github.com/lucasfrancaid/go-url-shortener/pkg/application/base"
	"github.com/lucasfrancaid/go-url-shortener/pkg/application/dto"
	"github.com/lucasfrancaid/go-url-shortener/pkg/port/repository"
)

type StatsUseCase struct {
	shortenerRepository repository.ShortenerRepository
}

func NewStatsUseCase(shortenerRepository repository.ShortenerRepository) StatsUseCase {
	return StatsUseCase{shortenerRepository: shortenerRepository}
}

func (u *StatsUseCase) Do(d dto.ShortenedDTO) (dto.ShortenerStatsDTO, error) {
	if err := u.validate(&d); err != nil {
		return dto.ShortenerStatsDTO{}, &base.Error{Type: base.VALIDATOR_ERROR, Err: err}
	}
	entity, err := u.shortenerRepository.Stats(d.ShortenedURL)
	if err != nil {
		if baseErr, ok := err.(*base.Error); ok {
			err = baseErr
		} else if strings.Contains(err.Error(), "not found") {
			err = &base.Error{Type: base.NOT_FOUND_ERROR, Err: err}
		} else {
			err = &base.Error{Type: base.INTERNAL_ERROR, Err: err}
		}
		return dto.ShortenerStatsDTO{}, err
	}
	return dto.ShortenerStatsDTO{Counter: entity.Counter}, nil
}

func (u *StatsUseCase) validate(d *dto.ShortenedDTO) error {
	if len(d.ShortenedURL) != 8 {
		return errors.New("HashedURL provided is invalid")
	}
	return nil
}

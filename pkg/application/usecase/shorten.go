package usecase

import (
	"encoding/hex"
	"errors"
	"strings"

	"crypto/md5"

	"github.com/lucasfrancaid/go-url-shortener/internal/pkg/infrastructure/config"
	base "github.com/lucasfrancaid/go-url-shortener/pkg/application/base"
	"github.com/lucasfrancaid/go-url-shortener/pkg/application/dto"
	"github.com/lucasfrancaid/go-url-shortener/pkg/domain"
	"github.com/lucasfrancaid/go-url-shortener/pkg/port/repository"
)

type ShortenUseCase struct {
	shortenerRepository repository.ShortenerRepository
}

func NewShortenUseCase(shortenerRepository repository.ShortenerRepository) ShortenUseCase {
	return ShortenUseCase{shortenerRepository: shortenerRepository}
}

func (u *ShortenUseCase) Do(d dto.ShortenDTO) (dto.ShortenedDTO, error) {
	err := u.validate(&d)
	if err != nil {
		return dto.ShortenedDTO{}, &base.Error{Type: base.VALIDATOR_ERROR, Err: err}
	}

	shortenedURL := u.short(d.URL)
	persisted, err := u.shortenerRepository.Exists(shortenedURL)
	if err == nil {
		if persisted.URL == d.URL {
			return u.toOutputDTO(persisted), nil
		}
		return dto.ShortenedDTO{}, &base.Error{
			Type: base.CONFLICT_ERROR,
			Err:  errors.New("conflict to try hash URL"),
		}
	}

	entity := domain.Shortener{HashedURL: shortenedURL, URL: d.URL}
	err = u.shortenerRepository.Add(entity)
	if err != nil {
		if baseErr, ok := err.(*base.Error); ok {
			err = baseErr
		} else {
			err = &base.Error{Type: base.INTERNAL_ERROR, Err: err}
		}
		return dto.ShortenedDTO{}, err
	}

	return u.toOutputDTO(entity), nil
}

func (u *ShortenUseCase) validate(d *dto.ShortenDTO) error {
	d.URL = strings.ToLower(strings.TrimSpace(d.URL))
	if len(d.URL) <= 10 {
		return errors.New("URL must to have more than 10 characters")
	}
	return nil
}

func (u *ShortenUseCase) short(URL string) string {
	md5Hash := md5.Sum([]byte(URL))
	shortenedURL := hex.EncodeToString(md5Hash[:4])
	return shortenedURL
}

func (u *ShortenUseCase) toOutputDTO(entity domain.Shortener) dto.ShortenedDTO {
	sURL := config.GetSettings().DOMAIN + "/u/" + entity.HashedURL
	return dto.ShortenedDTO{ShortenedURL: sURL}
}

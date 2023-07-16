package usecase

import (
	"testing"

	factory "github.com/lucasfrancaid/go-url-shortener/internal/pkg/infrastructure/factory/repository"
	"github.com/lucasfrancaid/go-url-shortener/pkg/application/dto"
	"github.com/lucasfrancaid/go-url-shortener/pkg/port/repository"
	"github.com/stretchr/testify/assert"
)

func TestNewShortenUseCase(t *testing.T) {
	r := factory.NewShortenerRepository()
	u := NewShortenUseCase(r)

	assert.IsType(t, ShortenUseCase{}, u)
	assert.Implements(t, (*repository.ShortenerRepository)(nil), u.shortenerRepository)

}

func TestShortenUseCase_Do(t *testing.T) {
	r := factory.NewShortenerRepository()
	u := NewShortenUseCase(r)
	d := dto.ShortenDTO{URL: "https://lucasfrancaid.com.br"}

	res, err := u.Do(d)

	assert.Nil(t, err)
	assert.NotNil(t, res.ShortenedURL)
}

func TestShortenUseCase_Do_WhenInvalidUrlShouldReturnError(t *testing.T) {
	r := factory.NewShortenerRepository()
	u := NewShortenUseCase(r)
	d := dto.ShortenDTO{URL: "InvalidUrl"}

	_, err := u.Do(d)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "URL must to have more than 10 characters")
}

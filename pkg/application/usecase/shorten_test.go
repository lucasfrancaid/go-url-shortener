package usecase

import (
	"testing"

	adapter "github.com/lucasfrancaid/go-url-shortener/pkg/adapter/repository/in_memory"
	"github.com/lucasfrancaid/go-url-shortener/pkg/application/dto"
	"github.com/stretchr/testify/assert"
)

func TestNewShortenUseCase(t *testing.T) {
	r := adapter.NewShortenerRepositoryInMemory()
	u := NewShortenUseCase(r)

	assert.IsType(t, ShortenUseCase{}, u)
	assert.IsType(t, &adapter.ShortenerRepositoryInMemory{}, u.shortenerRepository)
}

func TestShortenUseCase_Do(t *testing.T) {
	r := adapter.NewShortenerRepositoryInMemory()
	u := NewShortenUseCase(r)
	d := dto.ShortenDTO{URL: "https://lucasfrancaid.com.br"}

	res, err := u.Do(d)

	assert.Nil(t, err)
	assert.NotNil(t, res.ShortenedURL)
}

func TestShortenUseCase_Do_WhenInvalidUrlShouldReturnError(t *testing.T) {
	r := adapter.NewShortenerRepositoryInMemory()
	u := NewShortenUseCase(r)
	d := dto.ShortenDTO{URL: "InvalidUrl"}

	_, err := u.Do(d)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "URL must to have more than 10 characters")
}

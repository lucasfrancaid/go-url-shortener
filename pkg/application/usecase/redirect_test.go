package usecase

import (
	"testing"

	adapter "github.com/lucasfrancaid/go-url-shortener/pkg/adapter/repository/in_memory"
	"github.com/lucasfrancaid/go-url-shortener/pkg/application/dto"
	"github.com/lucasfrancaid/go-url-shortener/pkg/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewRedirectUseCase(t *testing.T) {
	r := adapter.NewShortenerRepositoryInMemory()
	u := NewRedirectUseCase(r)

	assert.IsType(t, RedirectUseCase{}, u)
	assert.IsType(t, &adapter.ShortenerRepositoryInMemory{}, u.shortenerRepository)
}

func TestRedirectUseCase_Do(t *testing.T) {
	r := adapter.NewShortenerRepositoryInMemory()
	m := domain.Shortener{HashedURL: "abcdefgh", URL: "https://any.com"}
	r.Add(m)

	u := NewRedirectUseCase(r)
	d := dto.ShortenedDTO{ShortenedURL: m.HashedURL}

	res, err := u.Do(d)

	assert.Nil(t, err)
	assert.Equal(t, m.URL, res.URL)
}

func TestRedirectUseCase_Do_WhenInvalidHashedUrlShouldReturnError(t *testing.T) {
	r := adapter.NewShortenerRepositoryInMemory()

	u := NewRedirectUseCase(r)
	d := dto.ShortenedDTO{ShortenedURL: "invalid"}

	_, err := u.Do(d)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "HashedURL provided is invalid")
}

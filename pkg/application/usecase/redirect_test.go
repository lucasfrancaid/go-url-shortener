package usecase

import (
	"testing"

	factory "github.com/lucasfrancaid/go-url-shortener/internal/pkg/infrastructure/factory/repository"
	"github.com/lucasfrancaid/go-url-shortener/pkg/application/dto"
	usecase_test "github.com/lucasfrancaid/go-url-shortener/pkg/application/usecase/test"
	"github.com/lucasfrancaid/go-url-shortener/pkg/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewRedirectUseCase(t *testing.T) {
	teardownTest := usecase_test.SetupUseCaseTest(t, nil, nil)
	defer teardownTest(t)

	r := factory.NewShortenerRepository()
	sr := factory.NewShortenerStatsRepository()
	u := NewRedirectUseCase(r, sr)

	assert.IsType(t, RedirectUseCase{}, u)
}

func TestRedirectUseCase_Do(t *testing.T) {
	s := domain.Shortener{HashedURL: "abcdefgh", URL: "https://any.com"}
	teardownTest := usecase_test.SetupUseCaseTest(t, s, nil)
	defer teardownTest(t)

	r := factory.NewShortenerRepository()
	sr := factory.NewShortenerStatsRepository()
	u := NewRedirectUseCase(r, sr)
	d := dto.ShortenedDTO{ShortenedURL: s.HashedURL}

	res, err := u.Do(d)

	assert.Nil(t, err)
	assert.Equal(t, s.URL, res.URL)
}

func TestRedirectUseCase_Do_WhenInvalidShortnedUrlShouldReturnError(t *testing.T) {
	teardownTest := usecase_test.SetupUseCaseTest(t, nil, nil)
	defer teardownTest(t)

	r := factory.NewShortenerRepository()
	sr := factory.NewShortenerStatsRepository()
	u := NewRedirectUseCase(r, sr)
	d := dto.ShortenedDTO{ShortenedURL: "invalid"}

	_, err := u.Do(d)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "HashedURL provided is invalid")
}

func TestRedirectUseCase_Do_WhenShortenedUrlDoesNotExistShouldReturnError(t *testing.T) {
	teardownTest := usecase_test.SetupUseCaseTest(t, nil, nil)
	defer teardownTest(t)

	r := factory.NewShortenerRepository()
	sr := factory.NewShortenerStatsRepository()
	u := NewRedirectUseCase(r, sr)
	d := dto.ShortenedDTO{ShortenedURL: "zxzcxzos"}

	_, err := u.Do(d)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "HashedURL not found")
}

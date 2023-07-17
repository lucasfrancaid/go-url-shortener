package usecase

import (
	"testing"

	factory "github.com/lucasfrancaid/go-url-shortener/internal/pkg/infrastructure/factory/repository"
	"github.com/lucasfrancaid/go-url-shortener/pkg/application/dto"
	usecase_test "github.com/lucasfrancaid/go-url-shortener/pkg/application/usecase/test"
	"github.com/stretchr/testify/assert"
)

func TestNewShortenUseCase(t *testing.T) {
	teardownTest := usecase_test.SetupUseCaseTest(t, nil, nil)
	defer teardownTest(t)

	r := factory.NewShortenerRepository()
	sr := factory.NewShortenerStatsRepository()
	u := NewShortenUseCase(r, sr)

	assert.IsType(t, ShortenUseCase{}, u)
}

func TestShortenUseCase_Do(t *testing.T) {
	teardownTest := usecase_test.SetupUseCaseTest(t, nil, nil)
	defer teardownTest(t)

	r := factory.NewShortenerRepository()
	sr := factory.NewShortenerStatsRepository()
	u := NewShortenUseCase(r, sr)
	d := dto.ShortenDTO{URL: "https://lucasfrancaid.com.br"}

	res, err := u.Do(d)

	assert.Nil(t, err)
	assert.NotNil(t, res.ShortenedURL)
}

func TestShortenUseCase_Do_WhenSuccessThenGetStatsShouldExist(t *testing.T) {
	teardownTest := usecase_test.SetupUseCaseTest(t, nil, nil)
	defer teardownTest(t)

	r := factory.NewShortenerRepository()
	sr := factory.NewShortenerStatsRepository()
	u := NewShortenUseCase(r, sr)
	d := dto.ShortenDTO{URL: "https://github.com.br/lucasfrancaid"}

	res, err := u.Do(d)
	assert.Nil(t, err)
	assert.NotNil(t, res.ShortenedURL)

	stats, err := u.statsRepository.Get(res.ShortenedURL[len(res.ShortenedURL)-8:])
	assert.Nil(t, err)
	assert.Equal(t, int64(0), stats.Counter)
}

func TestShortenUseCase_Do_WhenInvalidUrlShouldReturnError(t *testing.T) {
	teardownTest := usecase_test.SetupUseCaseTest(t, nil, nil)
	defer teardownTest(t)

	r := factory.NewShortenerRepository()
	sr := factory.NewShortenerStatsRepository()
	u := NewShortenUseCase(r, sr)
	d := dto.ShortenDTO{URL: "InvalidUrl"}

	_, err := u.Do(d)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "URL must to have more than 10 characters")
}

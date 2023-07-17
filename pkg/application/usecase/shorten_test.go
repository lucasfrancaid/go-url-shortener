package usecase

import (
	"testing"

	"github.com/lucasfrancaid/go-url-shortener/pkg/application/dto"
	usecase_test "github.com/lucasfrancaid/go-url-shortener/pkg/application/usecase/test"
	"github.com/stretchr/testify/assert"
)

func TestNewShortenUseCase(t *testing.T) {
	teardownTest := usecase_test.SetupUseCaseTest(t, nil, nil)
	defer teardownTest(t)

	u := NewShortenUseCase()

	assert.IsType(t, ShortenUseCase{}, u)
}

func TestShortenUseCase_Do(t *testing.T) {
	teardownTest := usecase_test.SetupUseCaseTest(t, nil, nil)
	defer teardownTest(t)

	u := NewShortenUseCase()
	d := dto.ShortenDTO{URL: "https://lucasfrancaid.com.br"}

	res, err := u.Do(d)

	assert.Nil(t, err)
	assert.NotNil(t, res.ShortenedURL)
}

func TestShortenUseCase_Do_WhenSuccessThenGetStatsShouldExist(t *testing.T) {
	teardownTest := usecase_test.SetupUseCaseTest(t, nil, nil)
	defer teardownTest(t)

	u := NewShortenUseCase()
	d := dto.ShortenDTO{URL: "https://github.com.br/lucasfrancaid"}

	res, err := u.Do(d)
	assert.Nil(t, err)
	assert.NotNil(t, res.ShortenedURL)

	stats, err := u.StatsRepository.Get(res.ShortenedURL[len(res.ShortenedURL)-8:])
	assert.Nil(t, err)
	assert.Equal(t, int64(0), stats.Counter)
}

func TestShortenUseCase_Do_WhenInvalidUrlShouldReturnError(t *testing.T) {
	teardownTest := usecase_test.SetupUseCaseTest(t, nil, nil)
	defer teardownTest(t)

	u := NewShortenUseCase()
	d := dto.ShortenDTO{URL: "InvalidUrl"}

	_, err := u.Do(d)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "URL must to have more than 10 characters")
}

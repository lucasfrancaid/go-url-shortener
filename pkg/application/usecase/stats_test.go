package usecase

import (
	"testing"

	factory "github.com/lucasfrancaid/go-url-shortener/internal/pkg/infrastructure/factory/repository"
	"github.com/lucasfrancaid/go-url-shortener/pkg/application/dto"
	usecase_test "github.com/lucasfrancaid/go-url-shortener/pkg/application/usecase/test"
	"github.com/stretchr/testify/assert"
)

func TestNewStatsUseCase(t *testing.T) {
	teardownTest := usecase_test.SetupUseCaseTest(t, nil, nil)
	defer teardownTest(t)

	r := factory.NewShortenerStatsRepository()
	u := NewStatsUseCase(r)

	assert.IsType(t, StatsUseCase{}, u)
}

func TestStatsUseCase_Do(t *testing.T) {
	hashedURL := "abcdefgh"
	teardownTest := usecase_test.SetupUseCaseTest(t, nil, hashedURL)
	defer teardownTest(t)

	r := factory.NewShortenerStatsRepository()
	u := NewStatsUseCase(r)
	d := dto.ShortenedDTO{ShortenedURL: hashedURL}

	res, err := u.Do(d)

	assert.Nil(t, err)
	assert.GreaterOrEqual(t, res.Counter, int64(0))
}

func TestStatsUseCase_Do_WhenInvalidShortenedUrlShouldReturnError(t *testing.T) {
	teardownTest := usecase_test.SetupUseCaseTest(t, nil, nil)
	defer teardownTest(t)

	r := factory.NewShortenerStatsRepository()
	u := NewStatsUseCase(r)
	d := dto.ShortenedDTO{ShortenedURL: "invalid"}

	_, err := u.Do(d)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "HashedURL provided is invalid")
}

func TestStatsUseCase_Do_WhenShortenedUrlDoesNotExistShouldReturnError(t *testing.T) {
	teardownTest := usecase_test.SetupUseCaseTest(t, nil, nil)
	defer teardownTest(t)

	r := factory.NewShortenerStatsRepository()

	u := NewStatsUseCase(r)
	d := dto.ShortenedDTO{ShortenedURL: "zxzcxzos"}

	_, err := u.Do(d)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "HashedURL not found")
}

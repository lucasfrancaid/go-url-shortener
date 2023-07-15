package usecase

import (
	"testing"

	adapter "github.com/lucasfrancaid/go-url-shortener/pkg/adapter/repository/in_memory"
	"github.com/lucasfrancaid/go-url-shortener/pkg/application/dto"
	"github.com/lucasfrancaid/go-url-shortener/pkg/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewStatsUseCase(t *testing.T) {
	r := adapter.NewShortenerRepositoryInMemory()
	u := NewStatsUseCase(r)

	assert.IsType(t, StatsUseCase{}, u)
	assert.IsType(t, &adapter.ShortenerRepositoryInMemory{}, u.shortenerRepository)
}

func TestStatsUseCase_Do(t *testing.T) {
	r := adapter.NewShortenerRepositoryInMemory()
	m := domain.Shortener{HashedURL: "abcdefgh", URL: "https://any.com"}
	r.Add(m)
	r.Read(m.HashedURL)

	u := NewStatsUseCase(r)
	d := dto.ShortenedDTO{ShortenedURL: m.HashedURL}

	res, err := u.Do(d)

	assert.Nil(t, err)
	assert.GreaterOrEqual(t, res.Counter, int64(0))
}

func TestStatsUseCase_Do_WhenInvalidShortenedUrlShouldReturnError(t *testing.T) {
	r := adapter.NewShortenerRepositoryInMemory()

	u := NewStatsUseCase(r)
	d := dto.ShortenedDTO{ShortenedURL: "invalid"}

	_, err := u.Do(d)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "HashedURL provided is invalid")
}

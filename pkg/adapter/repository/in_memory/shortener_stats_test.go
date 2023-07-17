package adapter

import (
	"testing"

	"github.com/lucasfrancaid/go-url-shortener/pkg/port/repository"
	"github.com/stretchr/testify/assert"
)

func setupShortenerStatsRepositoryInMemoryTest(tb testing.TB, d any) func(tb testing.TB) {
	r := NewShortenerStatsRepositoryInMemory()

	if hashedURL, ok := d.(string); ok {
		err := r.Set(hashedURL)
		assert.Nil(tb, err)
	}

	return func(tb testing.TB) {
		r.Counter = make(map[string]int64)
	}
}

func TestNewShortenerStatsRepositoryInMemory(t *testing.T) {
	teardownTest := setupShortenerStatsRepositoryInMemoryTest(t, nil)
	defer teardownTest(t)

	r := NewShortenerStatsRepositoryInMemory()

	assert.IsType(t, &ShortenerStatsRepositoryInMemory{}, r)
	assert.NotNil(t, r.Counter)
}

func TestShortenerStatsRepositoryInMemory_Set_WhenSucessShouldReturnNil(t *testing.T) {
	teardownTest := setupShortenerStatsRepositoryInMemoryTest(t, nil)
	defer teardownTest(t)

	r := NewShortenerStatsRepositoryInMemory()

	err := r.Set("any123")

	assert.Nil(t, err)
}

func TestShortenerStatsRepositoryInMemory_Set_WhenSuccessThenGetShouldReturnDomainWithValues(t *testing.T) {
	teardownTest := setupShortenerStatsRepositoryInMemoryTest(t, nil)
	defer teardownTest(t)

	hashedURL := "any343242"
	r := NewShortenerStatsRepositoryInMemory()

	err := r.Set(hashedURL)
	assert.Nil(t, err)

	stats, err := r.Get(hashedURL)
	assert.Nil(t, err)
	assert.Equal(t, hashedURL, stats.HashedURL)
	assert.Equal(t, int64(0), stats.Counter)
}

func TestShortenerStatsRepositoryInMemory_Set_WhenAlreadyExistShouldReturNil(t *testing.T) {
	hashedURL := "any12y3"
	teardownTest := setupShortenerRepositoryInMemoryTest(t, hashedURL)
	defer teardownTest(t)

	r := NewShortenerStatsRepositoryInMemory()

	err := r.Set("any12y3")

	assert.Nil(t, err)
}

func TestShortenerStatsRepositoryInMemory_Get_WhenSucessShouldReturnDomainWithValues(t *testing.T) {
	hashedURL := "some123"
	teardownTest := setupShortenerStatsRepositoryInMemoryTest(t, hashedURL)
	defer teardownTest(t)

	r := NewShortenerStatsRepositoryInMemory()

	stats, err := r.Get(hashedURL)

	assert.Nil(t, err)
	assert.Equal(t, int64(0), stats.Counter)
	assert.Equal(t, hashedURL, stats.HashedURL)
}

func TestShortenerStatsRepositoryInMemory_Get_WhenDoesNotExistShouldReturnError(t *testing.T) {
	teardownTest := setupShortenerStatsRepositoryInMemoryTest(t, nil)
	defer teardownTest(t)

	r := NewShortenerStatsRepositoryInMemory()

	_, err := r.Get("UnknownHashedURL")

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, repository.ErrShortenerStatsRepositoryGetNotFound)
}

func TestShortenerStatsRepositoryInMemory_Increment_WhenSuccessShouldReturnNil(t *testing.T) {
	hashedURL := "any1234"
	teardownTest := setupShortenerStatsRepositoryInMemoryTest(t, hashedURL)
	defer teardownTest(t)

	r := NewShortenerStatsRepositoryInMemory()

	err := r.Increment(hashedURL)

	assert.Nil(t, err)
}

func TestShortenerStatsRepositoryInMemory_Increment_WhenSuccessThenGetShouldReturnValueGreaterThanZero(t *testing.T) {
	hashedURL := "any12345"
	teardownTest := setupShortenerStatsRepositoryInMemoryTest(t, hashedURL)
	defer teardownTest(t)

	r := NewShortenerStatsRepositoryInMemory()

	err := r.Increment(hashedURL)
	assert.Nil(t, err)

	stats, err := r.Get(hashedURL)
	assert.Nil(t, err)
	assert.Greater(t, stats.Counter, int64(0))
}

func TestShortenerStatsRepositoryInMemory_Increment_WhenDoesNotExistShouldReturnError(t *testing.T) {
	teardownTest := setupShortenerStatsRepositoryInMemoryTest(t, nil)
	defer teardownTest(t)

	r := NewShortenerStatsRepositoryInMemory()

	err := r.Increment("UnknownHashedURL")

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, repository.ErrShortenerStatsRepositoryIncrementNotFound)
}

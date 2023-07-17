package adapter

import (
	"testing"

	"github.com/lucasfrancaid/go-url-shortener/pkg/port/repository"
	"github.com/stretchr/testify/assert"
)

func setupShortenerStatsRepositoryMemcachedTest(tb testing.TB, d any) func(tb testing.TB) {
	r := NewShortenerStatsRepositoryMemcached()

	if hashedURL, ok := d.(string); ok {
		err := r.Set(hashedURL)
		assert.Nil(tb, err)
	}

	return func(tb testing.TB) {
		r.Mc.DeleteAll()
	}
}

func TestNewShortenerStatsRepositoryMemcached(t *testing.T) {
	teardownTest := setupShortenerStatsRepositoryMemcachedTest(t, nil)
	defer teardownTest(t)

	r := NewShortenerStatsRepositoryMemcached()

	assert.IsType(t, &ShortenerStatsRepositoryMemcached{}, r)
	assert.NotNil(t, r.Mc)
}

func TestShortenerStatsRepositoryMemcached_Set_WhenSucessShouldReturnNil(t *testing.T) {
	teardownTest := setupShortenerStatsRepositoryMemcachedTest(t, nil)
	defer teardownTest(t)

	r := NewShortenerStatsRepositoryMemcached()

	err := r.Set("any123")

	assert.Nil(t, err)
}

func TestShortenerStatsRepositoryMemcached_Set_WhenSuccessThenGetShouldReturnDomainWithValues(t *testing.T) {
	teardownTest := setupShortenerStatsRepositoryMemcachedTest(t, nil)
	defer teardownTest(t)

	hashedURL := "any343242"
	r := NewShortenerStatsRepositoryMemcached()

	err := r.Set(hashedURL)
	assert.Nil(t, err)

	stats, err := r.Get(hashedURL)
	assert.Nil(t, err)
	assert.Equal(t, hashedURL, stats.HashedURL)
	assert.Equal(t, int64(0), stats.Counter)
}

func TestShortenerStatsRepositoryMemcached_Set_WhenAlreadyExistShouldReturNil(t *testing.T) {
	hashedURL := "any12y3"
	teardownTest := setupShortenerStatsRepositoryMemcachedTest(t, hashedURL)
	defer teardownTest(t)

	r := NewShortenerStatsRepositoryMemcached()

	err := r.Set("any12y3")

	assert.Nil(t, err)
}

func TestShortenerStatsRepositoryMemcached_Get_WhenSucessShouldReturnDomainWithValues(t *testing.T) {
	hashedURL := "some123"
	teardownTest := setupShortenerStatsRepositoryMemcachedTest(t, hashedURL)
	defer teardownTest(t)

	r := NewShortenerStatsRepositoryMemcached()

	stats, err := r.Get(hashedURL)

	assert.Nil(t, err)
	assert.Equal(t, int64(0), stats.Counter)
	assert.Equal(t, hashedURL, stats.HashedURL)
}

func TestShortenerStatsRepositoryMemcached_Get_WhenDoesNotExistShouldReturnError(t *testing.T) {
	teardownTest := setupShortenerStatsRepositoryMemcachedTest(t, nil)
	defer teardownTest(t)

	r := NewShortenerStatsRepositoryMemcached()

	_, err := r.Get("UnknownHashedURL")

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, repository.ErrShortenerStatsRepositoryGetNotFound)
}

func TestShortenerStatsRepositoryMemcached_Increment_WhenSuccessShouldReturnNil(t *testing.T) {
	hashedURL := "any1234"
	teardownTest := setupShortenerStatsRepositoryMemcachedTest(t, hashedURL)
	defer teardownTest(t)

	r := NewShortenerStatsRepositoryMemcached()

	err := r.Increment(hashedURL)

	assert.Nil(t, err)
}

func TestShortenerStatsRepositoryMemcached_Increment_WhenSuccessThenGetShouldReturnValueGreaterThanZero(t *testing.T) {
	hashedURL := "any12345"
	teardownTest := setupShortenerStatsRepositoryMemcachedTest(t, hashedURL)
	defer teardownTest(t)

	r := NewShortenerStatsRepositoryMemcached()

	err := r.Increment(hashedURL)
	assert.Nil(t, err)

	stats, err := r.Get(hashedURL)
	assert.Nil(t, err)
	assert.Greater(t, stats.Counter, int64(0))
}

func TestShortenerStatsRepositoryMemcached_Increment_WhenDoesNotExistShouldReturnError(t *testing.T) {
	teardownTest := setupShortenerStatsRepositoryMemcachedTest(t, nil)
	defer teardownTest(t)

	r := NewShortenerStatsRepositoryMemcached()

	err := r.Increment("UnknownHashedURL")

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, repository.ErrShortenerStatsRepositoryIncrementNotFound)
}

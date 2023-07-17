package adapter

import (
	"testing"

	"github.com/lucasfrancaid/go-url-shortener/pkg/port/repository"
	"github.com/stretchr/testify/assert"
)

func setupShortenerStatsRepositoryRedisTest(tb testing.TB, d any) func(tb testing.TB) {
	r := NewShortenerStatsRepositoryRedis()

	if hashedURL, ok := d.(string); ok {
		err := r.Set(hashedURL)
		assert.Nil(tb, err)
	}

	return func(tb testing.TB) {
		r.Rdb.FlushAll(r.Ctx)
	}
}

func TestNewShortenerStatsRepositoryRedis(t *testing.T) {
	teardownTest := setupShortenerStatsRepositoryRedisTest(t, nil)
	defer teardownTest(t)

	r := NewShortenerStatsRepositoryRedis()

	assert.IsType(t, &ShortenerStatsRepositoryRedis{}, r)
	assert.NotNil(t, r.Ctx)
	assert.NotNil(t, r.Rdb)
}

func TestShortenerStatsRepositoryRedis_Set_WhenSucessShouldReturnNil(t *testing.T) {
	teardownTest := setupShortenerStatsRepositoryRedisTest(t, nil)
	defer teardownTest(t)

	r := NewShortenerStatsRepositoryRedis()

	err := r.Set("any123")

	assert.Nil(t, err)
}

func TestShortenerStatsRepositoryRedis_Set_WhenSuccessThenGetShouldReturnDomainWithValues(t *testing.T) {
	teardownTest := setupShortenerStatsRepositoryRedisTest(t, nil)
	defer teardownTest(t)

	hashedURL := "any343242"
	r := NewShortenerStatsRepositoryRedis()

	err := r.Set(hashedURL)
	assert.Nil(t, err)

	stats, err := r.Get(hashedURL)
	assert.Nil(t, err)
	assert.Equal(t, hashedURL, stats.HashedURL)
	assert.Equal(t, int64(0), stats.Counter)
}

func TestShortenerStatsRepositoryRedis_Set_WhenAlreadyExistShouldReturNil(t *testing.T) {
	hashedURL := "any12y3"
	teardownTest := setupShortenerStatsRepositoryRedisTest(t, hashedURL)
	defer teardownTest(t)

	r := NewShortenerStatsRepositoryRedis()

	err := r.Set("any12y3")

	assert.Nil(t, err)
}

func TestShortenerStatsRepositoryRedis_Get_WhenSucessShouldReturnDomainWithValues(t *testing.T) {
	hashedURL := "some123"
	teardownTest := setupShortenerStatsRepositoryRedisTest(t, hashedURL)
	defer teardownTest(t)

	r := NewShortenerStatsRepositoryRedis()

	stats, err := r.Get(hashedURL)

	assert.Nil(t, err)
	assert.Equal(t, int64(0), stats.Counter)
	assert.Equal(t, hashedURL, stats.HashedURL)
}

func TestShortenerStatsRepositoryRedis_Get_WhenDoesNotExistShouldReturnError(t *testing.T) {
	teardownTest := setupShortenerStatsRepositoryRedisTest(t, nil)
	defer teardownTest(t)

	r := NewShortenerStatsRepositoryRedis()

	_, err := r.Get("UnknownHashedURL")

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, repository.ErrShortenerStatsRepositoryGetNotFound)
}

func TestShortenerStatsRepositoryRedis_Increment_WhenSuccessShouldReturnNil(t *testing.T) {
	hashedURL := "any1234"
	teardownTest := setupShortenerStatsRepositoryRedisTest(t, hashedURL)
	defer teardownTest(t)

	r := NewShortenerStatsRepositoryRedis()

	err := r.Increment(hashedURL)

	assert.Nil(t, err)
}

func TestShortenerStatsRepositoryRedis_Increment_WhenSuccessThenGetShouldReturnValueGreaterThanZero(t *testing.T) {
	hashedURL := "any12345"
	teardownTest := setupShortenerStatsRepositoryRedisTest(t, hashedURL)
	defer teardownTest(t)

	r := NewShortenerStatsRepositoryRedis()

	err := r.Increment(hashedURL)
	assert.Nil(t, err)

	stats, err := r.Get(hashedURL)
	assert.Nil(t, err)
	assert.Greater(t, stats.Counter, int64(0))
}

func TestShortenerStatsRepositoryRedis_Increment_WhenDoesNotExistShouldReturnError(t *testing.T) {
	teardownTest := setupShortenerStatsRepositoryRedisTest(t, nil)
	defer teardownTest(t)

	r := NewShortenerStatsRepositoryRedis()

	err := r.Increment("UnknownHashedURL1234567")

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, repository.ErrShortenerStatsRepositoryIncrementNotFound)
}

package adapter

import (
	"testing"

	"github.com/lucasfrancaid/go-url-shortener/internal/pkg/infrastructure/config"
	"github.com/lucasfrancaid/go-url-shortener/pkg/domain"
	"github.com/lucasfrancaid/go-url-shortener/pkg/port/repository"
	"github.com/stretchr/testify/assert"
)

func setupShortenerRepositoryRedisTest(tb testing.TB, d any) func(tb testing.TB) {
	if config.GetSettings().TEST_REPOSITORY_ADAPTER != config.RedisAdapter {
		tb.Skip()
	}

	r := NewShortenerRepositoryRedis()

	if data, ok := d.(domain.Shortener); ok {
		err := r.Add(data)
		assert.Nil(tb, err)
	}

	return func(tb testing.TB) {
		if data, ok := d.(domain.Shortener); ok {
			r.Rdb.Del(r.Ctx, data.HashedURL).Err()
		}
	}
}

func TestNewShortenerRepositoryRedis(t *testing.T) {
	teardownTest := setupShortenerRepositoryRedisTest(t, nil)
	defer teardownTest(t)

	r := NewShortenerRepositoryRedis()

	assert.IsType(t, &ShortenerRepositoryRedis{}, r)
	assert.NotNil(t, r.Ctx)
	assert.NotNil(t, r.Rdb)
}

func TestShortenerRepositoryRedis_Add_WhenSuccessShouldReturnNil(t *testing.T) {
	teardownTest := setupShortenerRepositoryRedisTest(t, nil)
	defer teardownTest(t)

	r := NewShortenerRepositoryRedis()
	d := domain.Shortener{HashedURL: "any", URL: "yna"}

	err := r.Add(d)

	assert.Nil(t, err)
}

func TestShortenerRepositoryRedis_Add_WhenSuccessThenReadShouldReturnDomainWithValues(t *testing.T) {
	teardownTest := setupShortenerRepositoryRedisTest(t, nil)
	defer teardownTest(t)

	r := NewShortenerRepositoryRedis()
	d := domain.Shortener{HashedURL: "any", URL: "yna"}

	err := r.Add(d)
	assert.Nil(t, err)

	data, err := r.Read(d.HashedURL)
	assert.Nil(t, err)
	assert.Equal(t, d, data)
}

func TestShortenerRepositoryRedis_Add_WhenAlreadyExistsShouldReturnNil(t *testing.T) {
	d := domain.Shortener{HashedURL: "any", URL: "yna"}
	teardownTest := setupShortenerRepositoryRedisTest(t, d)
	defer teardownTest(t)

	r := NewShortenerRepositoryRedis()

	err := r.Add(d)

	assert.Nil(t, err)
}

func TestShortenerRepositoryRedis_Read_WhenSuccessShouldReturnDomainWithValues(t *testing.T) {
	d := domain.Shortener{HashedURL: "xpto", URL: "https://xpto.com"}
	teardownTest := setupShortenerRepositoryRedisTest(t, d)
	defer teardownTest(t)

	r := NewShortenerRepositoryRedis()

	data, err := r.Read(d.HashedURL)

	assert.Nil(t, err)
	assert.Equal(t, d, data)
}

func TestShortenerRepositoryRedis_Read_WhenDoesNotExistShouldReturnError(t *testing.T) {
	teardownTest := setupShortenerRepositoryRedisTest(t, nil)
	defer teardownTest(t)

	r := NewShortenerRepositoryRedis()

	_, err := r.Read("UnknownHashedURL")

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, repository.ErrShortenerRepositoryReadNotFound)
}

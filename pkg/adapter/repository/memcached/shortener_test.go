package adapter

import (
	"testing"

	"github.com/lucasfrancaid/go-url-shortener/pkg/domain"
	"github.com/lucasfrancaid/go-url-shortener/pkg/port/repository"
	"github.com/stretchr/testify/assert"
)

func setupShortenerRepositoryMemcachedTest(tb testing.TB, d any) func(tb testing.TB) {
	r := NewShortenerRepositoryMemcached()

	if data, ok := d.(domain.Shortener); ok {
		err := r.Add(data)
		assert.Nil(tb, err)
	}

	return func(tb testing.TB) {
		r.Mc.DeleteAll()
	}
}

func TestNewShortenerRepositoryMemcached(t *testing.T) {
	teardownTest := setupShortenerRepositoryMemcachedTest(t, nil)
	defer teardownTest(t)

	r := NewShortenerRepositoryMemcached()

	assert.IsType(t, &ShortenerRepositoryMemcached{}, r)
	assert.NotNil(t, r.Mc)
}

func TestShortenerRepositoryMemcached_Add_WhenSuccessShouldReturnNil(t *testing.T) {
	teardownTest := setupShortenerRepositoryMemcachedTest(t, nil)
	defer teardownTest(t)

	r := NewShortenerRepositoryMemcached()
	d := domain.Shortener{HashedURL: "any", URL: "yna"}

	err := r.Add(d)

	assert.Nil(t, err)
}

func TestShortenerRepositoryMemcached_Add_WhenSuccessThenReadShouldReturnDomainWithValues(t *testing.T) {
	teardownTest := setupShortenerRepositoryMemcachedTest(t, nil)
	defer teardownTest(t)

	r := NewShortenerRepositoryMemcached()
	d := domain.Shortener{HashedURL: "any", URL: "yna"}

	err := r.Add(d)
	assert.Nil(t, err)

	data, err := r.Read(d.HashedURL)
	assert.Nil(t, err)
	assert.Equal(t, d, data)
}

func TestShortenerRepositoryMemcached_Add_WhenAlreadyExistsShouldReturnNil(t *testing.T) {
	d := domain.Shortener{HashedURL: "any", URL: "yna"}
	teardownTest := setupShortenerRepositoryMemcachedTest(t, d)
	defer teardownTest(t)

	r := NewShortenerRepositoryMemcached()

	err := r.Add(d)

	assert.Nil(t, err)
}

func TestShortenerRepositoryMemcached_Read_WhenSuccessShouldReturnDomainWithValues(t *testing.T) {
	d := domain.Shortener{HashedURL: "xpto", URL: "https://xpto.com"}
	teardownTest := setupShortenerRepositoryMemcachedTest(t, d)
	defer teardownTest(t)

	r := NewShortenerRepositoryMemcached()

	data, err := r.Read(d.HashedURL)

	assert.Nil(t, err)
	assert.Equal(t, d, data)
}

func TestShortenerRepositoryMemcached_Read_WhenDoesNotExistShouldReturnError(t *testing.T) {
	teardownTest := setupShortenerRepositoryMemcachedTest(t, nil)
	defer teardownTest(t)

	r := NewShortenerRepositoryMemcached()

	_, err := r.Read("UnknownHashedURL")

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, repository.ErrShortenerRepositoryReadNotFound)
}

package adapter

import (
	"testing"

	"github.com/lucasfrancaid/go-url-shortener/pkg/domain"
	"github.com/lucasfrancaid/go-url-shortener/pkg/port/repository"
	"github.com/stretchr/testify/assert"
)

func setupShortenerRepositoryInMemoryTest(tb testing.TB, d any) func(tb testing.TB) {
	r := NewShortenerRepositoryInMemory()

	if data, ok := d.(domain.Shortener); ok {
		err := r.Add(data)
		assert.Nil(tb, err)
	}

	return func(tb testing.TB) {
		r.Storage = make(map[string]domain.Shortener)
	}
}

func TestNewShortenerRepositoryInMemory(t *testing.T) {
	teardownTest := setupShortenerRepositoryInMemoryTest(t, nil)
	defer teardownTest(t)

	r := NewShortenerRepositoryInMemory()

	assert.IsType(t, &ShortenerRepositoryInMemory{}, r)
	assert.NotNil(t, r.Storage)
}

func TestShortenerRepositoryInMemory_Add_WhenSuccessShouldReturnNil(t *testing.T) {
	teardownTest := setupShortenerRepositoryInMemoryTest(t, nil)
	defer teardownTest(t)

	r := NewShortenerRepositoryInMemory()
	e := domain.Shortener{HashedURL: "any", URL: "yna"}

	err := r.Add(e)

	assert.Nil(t, err)
}

func TestShortenerRepositoryInMemory_Add_WhenSuccessThenReadShouldReturnDomainWithValues(t *testing.T) {
	teardownTest := setupShortenerRepositoryInMemoryTest(t, nil)
	defer teardownTest(t)

	r := NewShortenerRepositoryInMemory()
	d := domain.Shortener{HashedURL: "any", URL: "yna"}

	err := r.Add(d)
	assert.Nil(t, err)

	data, err := r.Read(d.HashedURL)
	assert.Nil(t, err)
	assert.Equal(t, d, data)
}

func TestShortenerRepositoryInMemory_Add_WhenAlreadyExistsShouldReturnNil(t *testing.T) {
	d := domain.Shortener{HashedURL: "any", URL: "yna"}
	teardownTest := setupShortenerRepositoryInMemoryTest(t, d)
	defer teardownTest(t)

	r := NewShortenerRepositoryInMemory()

	err := r.Add(d)

	assert.Nil(t, err)
}

func TestShortenerRepositoryInMemory_Read_WhenSuccessShouldReturnDomainWithValues(t *testing.T) {
	d := domain.Shortener{HashedURL: "xpto", URL: "https://xpto.com"}
	teardownTest := setupShortenerRepositoryInMemoryTest(t, d)
	defer teardownTest(t)

	r := NewShortenerRepositoryInMemory()

	data, err := r.Read(d.HashedURL)

	assert.Nil(t, err)
	assert.Equal(t, d, data)
}

func TestShortenerRepositoryInMemory_Read_WhenDoesNotExistShouldReturnError(t *testing.T) {
	teardownTest := setupShortenerRepositoryInMemoryTest(t, nil)
	defer teardownTest(t)

	r := NewShortenerRepositoryInMemory()

	_, err := r.Read("UnknownHashedURL")

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, repository.ErrShortenerRepositoryReadNotFound)
}

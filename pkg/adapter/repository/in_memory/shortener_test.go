package adapter

import (
	"testing"

	"github.com/lucasfrancaid/go-url-shortener/pkg/domain"
	"github.com/stretchr/testify/assert"
)

func TestNewShortenerRepositoryInMemory(t *testing.T) {
	r := NewShortenerRepositoryInMemory()

	assert.IsType(t, &ShortenerRepositoryInMemory{}, r)
	assert.NotNil(t, r.storage)
	assert.NotNil(t, r.counter)
}

func TestShortenerRepositoryInMemory_Add(t *testing.T) {
	r := NewShortenerRepositoryInMemory()
	e := domain.Shortener{HashedURL: "any", URL: "yna"}

	err := r.Add(e)

	assert.Nil(t, err)
	assert.Equal(t, r.storage[e.HashedURL], e)
}

func TestShortenerRepositoryInMemory_Add_WhenAlreadyExistsShouldReturnError(t *testing.T) {
	r := NewShortenerRepositoryInMemory()
	e := domain.Shortener{HashedURL: "any", URL: "yna"}
	r.Add(e)

	err := r.Add(e)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "HashedURL already exists")
}

func TestShortenerRepositoryInMemory_Read(t *testing.T) {
	r := NewShortenerRepositoryInMemory()
	e := domain.Shortener{HashedURL: "any", URL: "yna"}
	r.storage[e.HashedURL] = e

	entity, err := r.Read(e.HashedURL)

	assert.Nil(t, err)
	assert.Equal(t, e, entity)
}

func TestShortenerRepositoryInMemory_Read_WhenDoesNotExistShouldReturnError(t *testing.T) {
	r := NewShortenerRepositoryInMemory()

	_, err := r.Read("UnknownHashedURL")

	assert.NotNil(t, err)
	assert.EqualError(t, err, "HashedURL not found")
}

func TestShortenerRepositoryInMemory_Stats(t *testing.T) {
	r := NewShortenerRepositoryInMemory()
	e := domain.Shortener{HashedURL: "some", URL: "emos"}
	err := r.Add(e)
	assert.Nil(t, err)

	statsBeforeRead, err := r.Stats(e.HashedURL)

	assert.Nil(t, err)
	assert.NotNil(t, statsBeforeRead.Counter)
	assert.Equal(t, e.HashedURL, statsBeforeRead.HashedURL)

	r.Read(e.HashedURL)

	statsAfterRead, err := r.Stats(e.HashedURL)

	assert.Nil(t, err)
	assert.Equal(t, e.HashedURL, statsAfterRead.HashedURL)
	assert.GreaterOrEqual(t, statsAfterRead.Counter, statsBeforeRead.Counter)
}

func TestShortenerRepositoryInMemory_Stats_WhenDoesNotExistShouldReturnError(t *testing.T) {
	r := NewShortenerRepositoryInMemory()

	_, err := r.Stats("UnknownHashedURL")

	assert.NotNil(t, err)
	assert.EqualError(t, err, "HashedURL not found")
}

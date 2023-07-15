package controller

import (
	"testing"

	adapter "github.com/lucasfrancaid/go-url-shortener/pkg/adapter/repository/in_memory"
	"github.com/lucasfrancaid/go-url-shortener/pkg/application/dto"
	"github.com/lucasfrancaid/go-url-shortener/pkg/domain"
	"github.com/lucasfrancaid/go-url-shortener/pkg/port/presenter"
	"github.com/stretchr/testify/assert"
)

func TestNewShortenerController(t *testing.T) {
	r := adapter.NewShortenerRepositoryInMemory()
	c := NewShortenerController(r)

	assert.IsType(t, ShortenerController{}, c)
}

func TestShortenerController_Shorten(t *testing.T) {
	r := adapter.NewShortenerRepositoryInMemory()
	c := NewShortenerController(r)
	d := dto.ShortenDTO{URL: "https://github.com/lucasfrancaid"}

	p := c.Shorten(d)

	assert.Nil(t, p.Error)
	assert.NotNil(t, p.Data)
	assert.Equal(t, presenter.SUCCESS_CODE, p.StatusCode)
}

func TestShortenerController_Shorten_WhenInvalidUrlShouldReturnErrorAndValidatorErrorStatusCode(t *testing.T) {
	r := adapter.NewShortenerRepositoryInMemory()
	c := NewShortenerController(r)
	d := dto.ShortenDTO{URL: "xxx"}

	p := c.Shorten(d)

	assert.NotNil(t, p.Error)
	assert.Nil(t, p.Data)
	assert.Equal(t, presenter.VALIDATION_ERROR_CODE, p.StatusCode)
}

func TestShortenerController_Redirect(t *testing.T) {
	r := adapter.NewShortenerRepositoryInMemory()
	m := domain.Shortener{HashedURL: "abcdefgh", URL: "https://any.com"}
	r.Add(m)

	c := NewShortenerController(r)
	d := dto.ShortenedDTO{ShortenedURL: m.HashedURL}

	p := c.Redirect(d)

	assert.Nil(t, p.Error)
	assert.NotNil(t, p.Data)
	assert.Equal(t, presenter.REDIRECT_CODE, p.StatusCode)
}

func TestShortenerController_Redirect_WhenInvalidHashedUrlShouldReturnErrorAndValidatorErrorStatusCode(t *testing.T) {
	r := adapter.NewShortenerRepositoryInMemory()
	c := NewShortenerController(r)
	d := dto.ShortenedDTO{ShortenedURL: "xxx"}

	p := c.Redirect(d)

	assert.NotNil(t, p.Error)
	assert.Nil(t, p.Data)
	assert.Equal(t, presenter.VALIDATION_ERROR_CODE, p.StatusCode)
}

func TestShortenerController_Redirect_WhenDoesNotExistShouldReturnErrorAndNotFoundErrorStatusCode(t *testing.T) {
	r := adapter.NewShortenerRepositoryInMemory()
	c := NewShortenerController(r)
	d := dto.ShortenedDTO{ShortenedURL: "validurl"}

	p := c.Redirect(d)

	assert.NotNil(t, p.Error)
	assert.Nil(t, p.Data)
	assert.Equal(t, presenter.NOT_FOUND_ERROR_CODE, p.StatusCode)
}

func TestShortenerController_Stats(t *testing.T) {
	r := adapter.NewShortenerRepositoryInMemory()
	m := domain.Shortener{HashedURL: "abcdefgh", URL: "https://any.com"}
	r.Add(m)

	c := NewShortenerController(r)
	d := dto.ShortenedDTO{ShortenedURL: m.HashedURL}

	p := c.Stats(d)

	assert.Nil(t, p.Error)
	assert.NotNil(t, p.Data)
	assert.Equal(t, presenter.REDIRECT_CODE, p.StatusCode)
}

func TestShortenerController_Stats_WhenInvalidShortenedUrlShouldReturnErrorAndValidatorErrorStatusCode(t *testing.T) {
	r := adapter.NewShortenerRepositoryInMemory()
	c := NewShortenerController(r)
	d := dto.ShortenedDTO{ShortenedURL: "xxx"}

	p := c.Stats(d)

	assert.NotNil(t, p.Error)
	assert.Nil(t, p.Data)
	assert.Equal(t, presenter.VALIDATION_ERROR_CODE, p.StatusCode)
}

func TestShortenerController_Stats_WhenDoesNotExistShouldReturnErrorAndNotFoundErrorStatusCode(t *testing.T) {
	r := adapter.NewShortenerRepositoryInMemory()
	c := NewShortenerController(r)
	d := dto.ShortenedDTO{ShortenedURL: "validurl"}

	p := c.Stats(d)

	assert.NotNil(t, p.Error)
	assert.Nil(t, p.Data)
	assert.Equal(t, presenter.NOT_FOUND_ERROR_CODE, p.StatusCode)
}

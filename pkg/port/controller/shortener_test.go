package controller

import (
	"testing"

	adapter "github.com/lucasfrancaid/go-url-shortener/pkg/adapter/repository/in_memory"
	"github.com/lucasfrancaid/go-url-shortener/pkg/application/dto"
	"github.com/lucasfrancaid/go-url-shortener/pkg/port/presenter"
	"github.com/stretchr/testify/assert"
)

func TestNewShortenerController(t *testing.T) {
	r := adapter.NewShortenerRepositoryInMemory()
	c := NewShortenerController(&r)

	assert.IsType(t, ShortenerController{}, c)
}

func TestShortenerController_Shorten(t *testing.T) {
	r := adapter.NewShortenerRepositoryInMemory()
	c := NewShortenerController(&r)
	d := dto.ShortenDTO{URL: "https://github.com/lucasfrancaid"}

	p := c.Shorten(d)

	assert.Nil(t, p.Error)
	assert.NotNil(t, p.Data)
	assert.Equal(t, presenter.SUCCESS_CODE, p.StatusCode)
}

func TestShortenerController_Shorten_WhenInvalidUrlShouldReturnErrorAndValidatorErrorStatusCode(t *testing.T) {
	r := adapter.NewShortenerRepositoryInMemory()
	c := NewShortenerController(&r)
	d := dto.ShortenDTO{URL: "xxx"}

	p := c.Shorten(d)

	assert.NotNil(t, p.Error)
	assert.Nil(t, p.Data)
	assert.Equal(t, presenter.VALIDATION_ERROR_CODE, p.StatusCode)
}

package controller

import (
	"testing"

	"github.com/lucasfrancaid/go-url-shortener/internal/pkg/infrastructure/config"
	factory "github.com/lucasfrancaid/go-url-shortener/internal/pkg/infrastructure/factory/repository"
	in_memory "github.com/lucasfrancaid/go-url-shortener/pkg/adapter/repository/in_memory"
	memcached "github.com/lucasfrancaid/go-url-shortener/pkg/adapter/repository/memcached"
	redis "github.com/lucasfrancaid/go-url-shortener/pkg/adapter/repository/redis"
	"github.com/lucasfrancaid/go-url-shortener/pkg/application/dto"
	"github.com/lucasfrancaid/go-url-shortener/pkg/domain"
	"github.com/lucasfrancaid/go-url-shortener/pkg/port/presenter"
	"github.com/stretchr/testify/assert"
)

func setupShortenerControllerTest(tb testing.TB, d any, hashedURL any) func(tb testing.TB) {
	oldValue := config.GetSettings().REPOSITORY_ADAPTER
	config.GetSettings().REPOSITORY_ADAPTER = config.GetSettings().TEST_REPOSITORY_ADAPTER

	sr := factory.NewShortenerRepository()
	ssr := factory.NewShortenerStatsRepository()

	if data, ok := d.(domain.Shortener); ok {
		err := sr.Add(data)
		assert.Nil(tb, err)
	}

	if url, ok := hashedURL.(string); ok {
		err := ssr.Set(url)
		assert.Nil(tb, err)
	}

	return func(tb testing.TB) {
		config.GetSettings().REPOSITORY_ADAPTER = oldValue

		switch _sr := sr.(type) {
		case *in_memory.ShortenerRepositoryInMemory:
			_sr.Storage = make(map[string]domain.Shortener)
		case *memcached.ShortenerRepositoryMemcached:
			if data, ok := d.(domain.Shortener); ok {
				_sr.Mc.Delete(data.HashedURL)
			}
		case *redis.ShortenerRepositoryRedis:
			if data, ok := d.(domain.Shortener); ok {
				_sr.Rdb.Del(_sr.Ctx, data.HashedURL).Err()
			}
		}

		switch _ssr := ssr.(type) {
		case *in_memory.ShortenerStatsRepositoryInMemory:
			_ssr.Counter = make(map[string]int64)
		case *memcached.ShortenerStatsRepositoryMemcached:
			if url, ok := hashedURL.(string); ok {
				_ssr.Mc.Delete(url)
			}
		case *redis.ShortenerStatsRepositoryRedis:
			if url, ok := hashedURL.(string); ok {
				_ssr.Rdb.Del(_ssr.Ctx, url).Err()
			}
		}
	}
}

func TestNewShortenerController(t *testing.T) {
	teardownTest := setupShortenerControllerTest(t, nil, nil)
	defer teardownTest(t)

	c := NewShortenerController()

	assert.IsType(t, ShortenerController{}, c)
}

func TestShortenerController_Shorten_WhenValidDTOShouldReturnPresenterWithSuccessStatusCode(t *testing.T) {
	teardownTest := setupShortenerControllerTest(t, nil, nil)
	defer teardownTest(t)

	c := NewShortenerController()
	d := dto.ShortenDTO{URL: "https://github.com/lucasfrancaid"}

	p := c.Shorten(d)

	assert.IsType(t, presenter.Presenter{}, p)
	assert.Nil(t, p.Error)
	assert.NotNil(t, p.Data)
	assert.Equal(t, presenter.SUCCESS_CODE, p.StatusCode)
}

func TestShortenerController_Shorten_WhenInvalidDTOShouldReturnPresenterWithErrorAndValidatorErrorStatusCode(t *testing.T) {
	teardownTest := setupShortenerControllerTest(t, nil, nil)
	defer teardownTest(t)

	c := NewShortenerController()
	d := dto.ShortenDTO{URL: "xxx"}

	p := c.Shorten(d)

	assert.IsType(t, presenter.Presenter{}, p)
	assert.NotNil(t, p.Error)
	assert.Nil(t, p.Data)
	assert.Equal(t, presenter.VALIDATION_ERROR_CODE, p.StatusCode)
}

func TestShortenerController_Redirect_WhenValidDTOShouldReturnPresenterWithRedirectStatusCode(t *testing.T) {
	s := domain.Shortener{HashedURL: "abcdefgh", URL: "https://any.com"}
	teardownTest := setupShortenerControllerTest(t, s, nil)
	defer teardownTest(t)

	c := NewShortenerController()
	d := dto.ShortenedDTO{ShortenedURL: s.HashedURL}

	p := c.Redirect(d)

	assert.IsType(t, presenter.Presenter{}, p)
	assert.Nil(t, p.Error)
	assert.NotNil(t, p.Data)
	assert.Equal(t, presenter.REDIRECT_CODE, p.StatusCode)
}

func TestShortenerController_Redirect_WhenInvaliDTOShouldReturnPresenterWithErrorAndValidatorErrorStatusCode(t *testing.T) {
	teardownTest := setupShortenerControllerTest(t, nil, nil)
	defer teardownTest(t)

	c := NewShortenerController()
	d := dto.ShortenedDTO{ShortenedURL: "xxx"}

	p := c.Redirect(d)

	assert.IsType(t, presenter.Presenter{}, p)
	assert.NotNil(t, p.Error)
	assert.Nil(t, p.Data)
	assert.Equal(t, presenter.VALIDATION_ERROR_CODE, p.StatusCode)
}

func TestShortenerController_Redirect_WhenDoesNotExistShouldReturnPresenterWithErrorAndNotFoundErrorStatusCode(t *testing.T) {
	teardownTest := setupShortenerControllerTest(t, nil, nil)
	defer teardownTest(t)

	c := NewShortenerController()
	d := dto.ShortenedDTO{ShortenedURL: "validurl"}

	p := c.Redirect(d)

	assert.IsType(t, presenter.Presenter{}, p)
	assert.NotNil(t, p.Error)
	assert.Nil(t, p.Data)
	assert.Equal(t, presenter.NOT_FOUND_ERROR_CODE, p.StatusCode)
}

func TestShortenerController_Stats_WhenValidDTOShouldReturnPresenterWithSuccessStatusCode(t *testing.T) {
	hashedURL := "abcdefgh"
	s := domain.Shortener{HashedURL: hashedURL, URL: "https://any.com"}
	teardownTest := setupShortenerControllerTest(t, s, hashedURL)
	defer teardownTest(t)

	c := NewShortenerController()
	d := dto.ShortenedDTO{ShortenedURL: hashedURL}

	p := c.Stats(d)

	assert.IsType(t, presenter.Presenter{}, p)
	assert.Nil(t, p.Error)
	assert.NotNil(t, p.Data)
	assert.Equal(t, presenter.SUCCESS_CODE, p.StatusCode)
}

func TestShortenerController_Stats_WhenInvalidDTOShouldReturnPresenterWithErrorAndValidatorErrorStatusCode(t *testing.T) {
	teardownTest := setupShortenerControllerTest(t, nil, nil)
	defer teardownTest(t)

	c := NewShortenerController()
	d := dto.ShortenedDTO{ShortenedURL: "xxx"}

	p := c.Stats(d)

	assert.IsType(t, presenter.Presenter{}, p)
	assert.NotNil(t, p.Error)
	assert.Nil(t, p.Data)
	assert.Equal(t, presenter.VALIDATION_ERROR_CODE, p.StatusCode)
}

func TestShortenerController_Stats_WhenDoesNotExistShouldReturnPresenterWithErrorAndNotFoundErrorStatusCode(t *testing.T) {
	teardownTest := setupShortenerControllerTest(t, nil, nil)
	defer teardownTest(t)

	c := NewShortenerController()
	d := dto.ShortenedDTO{ShortenedURL: "validurl"}

	p := c.Stats(d)

	assert.IsType(t, presenter.Presenter{}, p)
	assert.NotNil(t, p.Error)
	assert.Nil(t, p.Data)
	assert.Equal(t, presenter.NOT_FOUND_ERROR_CODE, p.StatusCode)
}

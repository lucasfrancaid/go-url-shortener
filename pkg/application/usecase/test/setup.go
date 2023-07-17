package usecase_test

import (
	"testing"

	"github.com/lucasfrancaid/go-url-shortener/internal/pkg/infrastructure/config"
	factory "github.com/lucasfrancaid/go-url-shortener/internal/pkg/infrastructure/factory/repository"
	in_memory "github.com/lucasfrancaid/go-url-shortener/pkg/adapter/repository/in_memory"
	memcached "github.com/lucasfrancaid/go-url-shortener/pkg/adapter/repository/memcached"
	redis "github.com/lucasfrancaid/go-url-shortener/pkg/adapter/repository/redis"
	"github.com/lucasfrancaid/go-url-shortener/pkg/domain"
	"github.com/stretchr/testify/assert"
)

func SetupUseCaseTest(tb testing.TB, d any, hashedURL any) func(tb testing.TB) {
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

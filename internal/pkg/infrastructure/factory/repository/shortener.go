package factory

import (
	"github.com/lucasfrancaid/go-url-shortener/internal/pkg/infrastructure/config"
	in_memory "github.com/lucasfrancaid/go-url-shortener/pkg/adapter/repository/in_memory"
	memcached "github.com/lucasfrancaid/go-url-shortener/pkg/adapter/repository/memcached"
	redis "github.com/lucasfrancaid/go-url-shortener/pkg/adapter/repository/redis"
	"github.com/lucasfrancaid/go-url-shortener/pkg/port/repository"
)

func NewShortenerRepository() repository.ShortenerRepository {
	switch config.GetSettings().REPOSITORY_ADAPTER {
	case config.RedisAdapter:
		return redis.NewShortenerRepositoryRedis()
	case config.MemcachedAdapter:
		return memcached.NewShortenerRepositoryMemcached()
	default:
		return in_memory.NewShortenerRepositoryInMemory()
	}
}

func NewShortenerStatsRepository() repository.ShortenerStatsRepository {
	switch config.GetSettings().REPOSITORY_ADAPTER {
	case config.RedisAdapter:
		return redis.NewShortenerStatsRepositoryRedis()
	case config.MemcachedAdapter:
		return memcached.NewShortenerStatsRepositoryMemcached()
	default:
		return in_memory.NewShortenerStatsRepositoryInMemory()
	}
}

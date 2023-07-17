package factory

import (
	"github.com/lucasfrancaid/go-url-shortener/internal/pkg/infrastructure/config"
	in_memory "github.com/lucasfrancaid/go-url-shortener/pkg/adapter/repository/in_memory"
	memcached "github.com/lucasfrancaid/go-url-shortener/pkg/adapter/repository/memcached"
	"github.com/lucasfrancaid/go-url-shortener/pkg/port/repository"
)

func NewShortenerRepository() repository.ShortenerRepository {
	switch config.GetSettings().REPOSITORY_ADAPTER {
	case "memcached":
		return memcached.NewShortenerRepositoryMemcached()
	default:
		return in_memory.NewShortenerRepositoryInMemory()
	}
}

func NewShortenerStatsRepository() repository.ShortenerStatsRepository {
	switch config.GetSettings().REPOSITORY_ADAPTER {
	case "memcached":
		return memcached.NewShortenerStatsRepositoryMemcached()
	default:
		return in_memory.NewShortenerStatsRepositoryInMemory()
	}
}

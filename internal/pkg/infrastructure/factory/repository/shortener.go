package factory

import (
	"github.com/lucasfrancaid/go-url-shortener/internal/pkg/infrastructure"
	in_memory "github.com/lucasfrancaid/go-url-shortener/pkg/adapter/repository/in_memory"
	memcached "github.com/lucasfrancaid/go-url-shortener/pkg/adapter/repository/memcached"
	"github.com/lucasfrancaid/go-url-shortener/pkg/port/repository"
)

func NewShortenerRepository() repository.ShortenerRepository {
	settings := infrastructure.Settings()

	switch settings.REPOSITORY_ADAPTER {
	case "memcached":
		return memcached.NewShortenerRepositoryMemcached()
	default:
		return in_memory.NewShortenerRepositoryInMemory()
	}
}

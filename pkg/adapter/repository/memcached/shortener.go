package adapter

import (
	"strings"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/lucasfrancaid/go-url-shortener/internal/pkg/infrastructure/config"
	"github.com/lucasfrancaid/go-url-shortener/pkg/domain"
	"github.com/lucasfrancaid/go-url-shortener/pkg/port/repository"
)

type ShortenerRepositoryMemcached struct {
	Mc *memcache.Client
}

func NewShortenerRepositoryMemcached() *ShortenerRepositoryMemcached {
	return &ShortenerRepositoryMemcached{
		Mc: memcache.New(config.GetSettings().MEMCACHED_URL),
	}
}

func (r *ShortenerRepositoryMemcached) Add(entity domain.Shortener) error {
	exists, _ := r.Mc.Get(entity.HashedURL)
	if exists != nil {
		return nil
	}
	item := &memcache.Item{Key: entity.HashedURL, Value: []byte(entity.URL)}
	return r.Mc.Set(item)
}

func (r *ShortenerRepositoryMemcached) Read(HashedURL string) (domain.Shortener, error) {
	item, err := r.Mc.Get(HashedURL)
	if err != nil {
		if strings.Contains(err.Error(), "memcache: cache miss") {
			err = repository.ErrShortenerRepositoryReadNotFound
		}
		return domain.Shortener{}, err
	}

	return domain.Shortener{HashedURL: item.Key, URL: string(item.Value)}, nil
}

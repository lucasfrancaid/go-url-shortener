package adapter

import (
	"strconv"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/lucasfrancaid/go-url-shortener/internal/pkg/infrastructure/config"
	"github.com/lucasfrancaid/go-url-shortener/pkg/domain"
	"github.com/lucasfrancaid/go-url-shortener/pkg/port/repository"
)

type ShortenerStatsRepositoryMemcached struct {
	Mc *memcache.Client
}

func NewShortenerStatsRepositoryMemcached() *ShortenerStatsRepositoryMemcached {
	return &ShortenerStatsRepositoryMemcached{
		Mc: memcache.New(config.GetSettings().MEMCACHED_URL),
	}
}

func (r *ShortenerStatsRepositoryMemcached) buildKey(Key string) string {
	return Key + "_stats"
}

func (r *ShortenerStatsRepositoryMemcached) Set(HashedURL string) error {
	return r.Mc.Set(&memcache.Item{Key: r.buildKey(HashedURL), Value: []byte("0")})
}

func (r *ShortenerStatsRepositoryMemcached) Increment(HashedURL string) error {
	_, err := r.Mc.Increment(r.buildKey(HashedURL), 1)
	if err == memcache.ErrCacheMiss {
		err = repository.ErrShortenerStatsRepositoryIncrementNotFound
	}
	return err
}

func (r *ShortenerStatsRepositoryMemcached) Get(HashedURL string) (domain.ShortenerStats, error) {
	i, err := r.Mc.Get(r.buildKey(HashedURL))
	if err != nil {
		if err == memcache.ErrCacheMiss {
			err = repository.ErrShortenerStatsRepositoryGetNotFound
		}
		return domain.ShortenerStats{}, err
	}
	c, err := strconv.ParseInt(string(i.Value), 10, 64)
	if err != nil {
		return domain.ShortenerStats{}, err
	}
	return domain.ShortenerStats{HashedURL: HashedURL, Counter: c}, nil
}

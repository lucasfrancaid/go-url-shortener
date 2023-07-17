package adapter

import (
	"context"

	"github.com/lucasfrancaid/go-url-shortener/internal/pkg/infrastructure/config"
	"github.com/lucasfrancaid/go-url-shortener/pkg/domain"
	"github.com/lucasfrancaid/go-url-shortener/pkg/port/repository"
	"github.com/redis/go-redis/v9"
)

type ShortenerRepositoryRedis struct {
	Ctx context.Context
	Rdb *redis.Client
}

func NewShortenerRepositoryRedis() *ShortenerRepositoryRedis {
	settings := config.GetSettings()
	return &ShortenerRepositoryRedis{
		Ctx: context.Background(),
		Rdb: redis.NewClient(&redis.Options{
			Addr:     settings.REDIS_URL,
			Password: settings.REDIS_PASSWORD,
			DB:       settings.REDIS_DB,
		}),
	}
}

func (r *ShortenerRepositoryRedis) Add(entity domain.Shortener) error {
	_, err := r.Rdb.Get(r.Ctx, entity.HashedURL).Result()
	if err == nil {
		return nil
	}
	return r.Rdb.Set(r.Ctx, entity.HashedURL, entity.URL, 0).Err()
}

func (r *ShortenerRepositoryRedis) Read(HashedURL string) (domain.Shortener, error) {
	item, err := r.Rdb.Get(r.Ctx, HashedURL).Result()
	if err != nil {
		if err == redis.Nil {
			err = repository.ErrShortenerRepositoryReadNotFound
		}
		return domain.Shortener{}, err
	}

	return domain.Shortener{HashedURL: HashedURL, URL: item}, nil
}

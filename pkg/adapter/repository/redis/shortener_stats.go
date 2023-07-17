package adapter

import (
	"context"
	"strconv"

	"github.com/lucasfrancaid/go-url-shortener/internal/pkg/infrastructure/config"
	"github.com/lucasfrancaid/go-url-shortener/pkg/domain"
	"github.com/lucasfrancaid/go-url-shortener/pkg/port/repository"
	"github.com/redis/go-redis/v9"
)

type ShortenerStatsRepositoryRedis struct {
	Ctx context.Context
	Rdb *redis.Client
}

func NewShortenerStatsRepositoryRedis() *ShortenerStatsRepositoryRedis {
	settings := config.GetSettings()
	return &ShortenerStatsRepositoryRedis{
		Ctx: context.Background(),
		Rdb: redis.NewClient(&redis.Options{
			Addr:     settings.REDIS_URL,
			Password: settings.REDIS_PASSWORD,
			DB:       settings.REDIS_DB,
		}),
	}
}

func (r *ShortenerStatsRepositoryRedis) buildKey(Key string) string {
	return Key + "_stats"
}

func (r *ShortenerStatsRepositoryRedis) Set(HashedURL string) error {
	return r.Rdb.Set(r.Ctx, r.buildKey(HashedURL), 0, 0).Err()
}

func (r *ShortenerStatsRepositoryRedis) Increment(HashedURL string) error {
	_, err := r.Rdb.Get(r.Ctx, r.buildKey(HashedURL)).Result()
	if err != nil {
		if err == redis.Nil {
			err = repository.ErrShortenerStatsRepositoryIncrementNotFound
		}
		return err
	}
	_, err = r.Rdb.Incr(r.Ctx, r.buildKey(HashedURL)).Result()
	return err
}

func (r *ShortenerStatsRepositoryRedis) Get(HashedURL string) (domain.ShortenerStats, error) {
	item, err := r.Rdb.Get(r.Ctx, r.buildKey(HashedURL)).Result()
	if err != nil {
		if err == redis.Nil {
			err = repository.ErrShortenerStatsRepositoryGetNotFound
		}
		return domain.ShortenerStats{}, err
	}
	c, err := strconv.ParseInt(string(item), 10, 64)
	if err != nil {
		return domain.ShortenerStats{}, err
	}
	return domain.ShortenerStats{HashedURL: HashedURL, Counter: c}, nil
}

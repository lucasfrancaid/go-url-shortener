package adapter

import (
	"github.com/lucasfrancaid/go-url-shortener/pkg/domain"
	"github.com/lucasfrancaid/go-url-shortener/pkg/port/repository"
)

type ShortenerStatsRepositoryInMemory struct {
	Counter map[string]int64
}

var globalCounter map[string]int64 = make(map[string]int64)

func NewShortenerStatsRepositoryInMemory() *ShortenerStatsRepositoryInMemory {
	return &ShortenerStatsRepositoryInMemory{
		Counter: globalCounter,
	}
}

func (r *ShortenerStatsRepositoryInMemory) Set(HashedURL string) error {
	r.Counter[HashedURL] += 0
	return nil
}

func (r *ShortenerStatsRepositoryInMemory) Increment(HashedURL string) error {
	_, ok := r.Counter[HashedURL]
	if !ok {
		return repository.ErrShortenerStatsRepositoryIncrementNotFound
	}
	r.Counter[HashedURL] += 1
	return nil
}

func (r *ShortenerStatsRepositoryInMemory) Get(HashedURL string) (domain.ShortenerStats, error) {
	c, ok := r.Counter[HashedURL]
	if !ok {
		return domain.ShortenerStats{}, repository.ErrShortenerStatsRepositoryGetNotFound
	}
	return domain.ShortenerStats{HashedURL: HashedURL, Counter: c}, nil
}

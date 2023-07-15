package adapter

import (
	"errors"

	"github.com/lucasfrancaid/go-url-shortener/pkg/domain"
)

type ShortenerRepositoryInMemory struct {
	storage map[string]domain.Shortener
	counter map[string]int64
}

var globalStorage map[string]domain.Shortener = make(map[string]domain.Shortener)
var globalCounter map[string]int64 = make(map[string]int64)

func NewShortenerRepositoryInMemory() *ShortenerRepositoryInMemory {
	return &ShortenerRepositoryInMemory{
		storage: globalStorage,
		counter: globalCounter,
	}
}

func (r *ShortenerRepositoryInMemory) Add(entity domain.Shortener) error {
	_, ok := r.storage[entity.HashedURL]
	if ok {
		return errors.New("HashedURL already exists")
	}
	r.storage[entity.HashedURL] = entity
	r.counter[entity.HashedURL] += 0
	return nil
}

func (r *ShortenerRepositoryInMemory) Read(HashedURL string) (domain.Shortener, error) {
	entity, ok := r.storage[HashedURL]
	if !ok {
		return domain.Shortener{}, errors.New("HashedURL not found")
	}
	r.counter[HashedURL] += 1
	return entity, nil
}

func (r *ShortenerRepositoryInMemory) Stats(HashedURL string) (domain.ShortenerStats, error) {
	c, ok := r.counter[HashedURL]
	if !ok {
		return domain.ShortenerStats{}, errors.New("HashedURL not found")
	}
	return domain.ShortenerStats{HashedURL: HashedURL, Counter: c}, nil
}

func (r *ShortenerRepositoryInMemory) Exists(HashedURL string) (domain.Shortener, error) {
	entity, ok := r.storage[HashedURL]
	if !ok {
		return domain.Shortener{}, errors.New("HashedURL not found")
	}
	return entity, nil
}

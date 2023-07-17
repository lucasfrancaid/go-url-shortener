package adapter

import (
	"github.com/lucasfrancaid/go-url-shortener/pkg/domain"
	"github.com/lucasfrancaid/go-url-shortener/pkg/port/repository"
)

type ShortenerRepositoryInMemory struct {
	Storage map[string]domain.Shortener
}

var globalStorage map[string]domain.Shortener = make(map[string]domain.Shortener)

func NewShortenerRepositoryInMemory() *ShortenerRepositoryInMemory {
	return &ShortenerRepositoryInMemory{
		Storage: globalStorage,
	}
}

func (r *ShortenerRepositoryInMemory) Add(entity domain.Shortener) error {
	_, ok := r.Storage[entity.HashedURL]
	if ok {
		return nil
	}
	r.Storage[entity.HashedURL] = entity
	return nil
}

func (r *ShortenerRepositoryInMemory) Read(HashedURL string) (domain.Shortener, error) {
	entity, ok := r.Storage[HashedURL]
	if !ok {
		return domain.Shortener{}, repository.ErrShortenerRepositoryReadNotFound
	}
	return entity, nil
}

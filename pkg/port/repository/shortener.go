package repository

import (
	"errors"

	"github.com/lucasfrancaid/go-url-shortener/pkg/domain"
)

var (
	ErrShortenerRepositoryReadNotFound = errors.New("HashedURL not found")
)

type ShortenerRepository interface {
	Add(entity domain.Shortener) error
	Read(HashedURL string) (domain.Shortener, error)
}

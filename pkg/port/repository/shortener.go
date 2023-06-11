package repository

import "github.com/lucasfrancaid/go-url-shortener/pkg/domain"

type ShortenerRepository interface {
	Add(entity domain.Shortener) error
	Read(HashedURL string) (domain.Shortener, error)
	Stats(HashedURL string) (domain.ShortenerStats, error)
}

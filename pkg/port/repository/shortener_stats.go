package repository

import (
	"errors"

	"github.com/lucasfrancaid/go-url-shortener/pkg/domain"
)

var (
	ErrShortenerStatsRepositoryIncrementNotFound = errors.New("HashedURL not found")
	ErrShortenerStatsRepositoryGetNotFound       = errors.New("HashedURL not found")
)

type ShortenerStatsRepository interface {
	Set(HashedURL string) error
	Increment(HashedURL string) error
	Get(HashedURL string) (domain.ShortenerStats, error)
}

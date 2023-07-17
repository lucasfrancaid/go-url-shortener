package usecase_test

import (
	"testing"

	"github.com/lucasfrancaid/go-url-shortener/internal/pkg/infrastructure/config"
	factory "github.com/lucasfrancaid/go-url-shortener/internal/pkg/infrastructure/factory/repository"
	in_memory "github.com/lucasfrancaid/go-url-shortener/pkg/adapter/repository/in_memory"
	"github.com/lucasfrancaid/go-url-shortener/pkg/domain"
	"github.com/stretchr/testify/assert"
)

func SetupUseCaseTest(tb testing.TB, d any, hashedURL any) func(tb testing.TB) {
	oldValue := config.GetSettings().REPOSITORY_ADAPTER
	config.GetSettings().REPOSITORY_ADAPTER = "in_memory"

	r := factory.NewShortenerRepository()
	sr := factory.NewShortenerStatsRepository()

	if data, ok := d.(domain.Shortener); ok {
		err := r.Add(data)
		assert.Nil(tb, err)
	}

	if url, ok := hashedURL.(string); ok {
		err := sr.Set(url)
		assert.Nil(tb, err)
	}

	return func(tb testing.TB) {
		config.GetSettings().REPOSITORY_ADAPTER = oldValue
		if rInMemory, ok := r.(*in_memory.ShortenerRepositoryInMemory); ok {
			rInMemory.Storage = make(map[string]domain.Shortener)
		}
		if srInMemory, ok := sr.(*in_memory.ShortenerStatsRepositoryInMemory); ok {
			srInMemory.Counter = make(map[string]int64)

		}
	}
}

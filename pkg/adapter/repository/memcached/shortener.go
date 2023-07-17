package adapter

import (
	"bytes"
	"encoding/gob"
	"errors"
	"strings"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/lucasfrancaid/go-url-shortener/internal/pkg/infrastructure/config"
	"github.com/lucasfrancaid/go-url-shortener/pkg/domain"
)

type ShortenerRepositoryMemcached struct {
	mc *memcache.Client
}

type itemValue struct {
	URL     string
	Counter int64
}

func NewShortenerRepositoryMemcached() *ShortenerRepositoryMemcached {
	return &ShortenerRepositoryMemcached{
		mc: memcache.New(config.GetSettings().MEMCACHED_URL),
	}
}

func (r *ShortenerRepositoryMemcached) set(Key string, Value itemValue) error {
	var buffer bytes.Buffer
	enc := gob.NewEncoder(&buffer)
	err := enc.Encode(Value)
	if err != nil {
		return err
	}

	item := &memcache.Item{Key: Key, Value: []byte(buffer.Bytes())}
	err = r.mc.Set(item)
	if err != nil {
		return err
	}
	return nil
}

func (r *ShortenerRepositoryMemcached) get(Key string) (itemValue, error) {
	item, err := r.mc.Get(Key)
	if err != nil {
		if strings.Contains(err.Error(), "memcache: cache miss") {
			err = errors.New("HashedURL not found")
		}
		return itemValue{}, err
	}

	buffer := bytes.NewBuffer(item.Value)
	dec := gob.NewDecoder(buffer)

	var value itemValue
	err = dec.Decode(&value)
	if err != nil {
		return itemValue{}, err
	}
	return value, nil
}

func (r *ShortenerRepositoryMemcached) Add(entity domain.Shortener) error {
	exists, _ := r.mc.Get(entity.HashedURL)
	if exists != nil {
		return nil
	}
	value := itemValue{URL: entity.URL, Counter: int64(0)}
	err := r.set(entity.HashedURL, value)
	if err != nil {
		return err
	}
	return nil
}

func (r *ShortenerRepositoryMemcached) Read(HashedURL string) (domain.Shortener, error) {
	value, err := r.get(HashedURL)
	if err != nil {
		return domain.Shortener{}, err
	}

	err = r.set(HashedURL, itemValue{URL: value.URL, Counter: value.Counter + 1})
	if err != nil {
		return domain.Shortener{}, err
	}

	return domain.Shortener{HashedURL: HashedURL, URL: value.URL}, nil
}

func (r *ShortenerRepositoryMemcached) Stats(HashedURL string) (domain.ShortenerStats, error) {
	value, err := r.get(HashedURL)
	if err != nil {
		return domain.ShortenerStats{}, err
	}
	return domain.ShortenerStats{HashedURL: HashedURL, Counter: value.Counter}, nil
}

func (r *ShortenerRepositoryMemcached) Exists(HashedURL string) (domain.Shortener, error) {
	value, err := r.get(HashedURL)
	if err != nil {
		return domain.Shortener{}, err
	}
	return domain.Shortener{HashedURL: HashedURL, URL: value.URL}, nil
}

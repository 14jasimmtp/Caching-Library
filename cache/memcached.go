package cache

import (
	"time"

	"github.com/bradfitz/gomemcache/memcache"
)

type memcached struct {
	client *memcache.Client
}

func NewMemcachedCache(server string) *memcached {
	return &memcached{client: memcache.New(server)}
}

func (c *memcached) Set(key string, value []byte, ttl time.Duration) error {
	return c.client.Set(&memcache.Item{
		Key:        key,
		Value:      value,
		Expiration: int32(ttl.Seconds()),
	})
}

func (c *memcached) Get(key string) (any, error) {
	item, err := c.client.Get(key)
	if err != nil {
		return nil, err
	}
	return string(item.Value), nil
}

func (c *memcached) Delete(key string) error {
	return c.client.Delete(key)
}

func (c *memcached) Clear() {
	c.client.DeleteAll()
}

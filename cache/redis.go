package cache

import (
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/copier"
)

type rs struct {
	Client *redis.Client
}

type RedisOptions struct {
	Addr     string
	Network  string
	Username string
	DB       int
	Password string
}

func NewRedisCache(opt *RedisOptions) Methods {
	RO:=redis.Options{}
	copier.Copy(&RO,&opt)
	return &rs{Client: redis.NewClient(&RO)}
}

func (c *rs) Set(key string, value interface{}, ttl time.Duration) error {
	return c.Client.Set(c.Client.Context(), key, value, ttl).Err()
}

func (c *rs) Get(key string) (any, error) {
	return c.Client.Get(c.Client.Context(), key).Result()
}

func (c *rs) Delete(key string) error {
	return c.Client.Del(c.Client.Context(), key).Err()
}

func (c *rs) Clear() {
	c.Client.FlushDB(c.Client.Context())
}
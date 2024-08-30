package cache

import (
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/copier"
)

type Rs struct {
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
	return &Rs{Client: redis.NewClient(&RO)}
}

func (c *Rs) Set(key string, value interface{}, ttl time.Duration) error {
	return c.Client.Set(c.Client.Context(), key, value, ttl).Err()
}

func (c *Rs) Get(key string) (any, error) {
	return c.Client.Get(c.Client.Context(), key).Result()
}

func (c *Rs) Delete(key string) error {
	return c.Client.Del(c.Client.Context(), key).Err()
}

func (c *Rs) Clear() {
	c.Client.FlushDB(c.Client.Context())
}
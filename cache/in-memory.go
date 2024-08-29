package cache

import (
	"container/list"
	"errors"
	"sync"
	"time"
)

type inMemory struct {
	Size  int
	Cache map[string]*list.Element
	DDL   *list.List
	mu    sync.RWMutex
}

type node struct {
	Key   string
	Value any
	TTL   time.Time
}

func NewInMemoryCache(Size int) Methods {
	cache := &inMemory{
		Size:  Size,
		Cache: make(map[string]*list.Element),
		DDL:   list.New(),
	}

	go cache.ExpiryWorker()

	return cache

}

func (c *inMemory) Set(key string, value interface{}, ttl time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if elem, found := c.Cache[key]; found {
		elem.Value.(*node).Value = value
		elem.Value.(*node).TTL = time.Now().Add(ttl)
		c.DDL.MoveToFront(elem)
	} else {
		if c.DDL.Len() == c.Size {
			last := c.DDL.Back()
			delete(c.Cache, last.Value.(*node).Key)
			c.DDL.Remove(last)
		}
		node := &node{Key: key, Value: value, TTL: time.Now().Add(ttl)}
		element := c.DDL.PushFront(node)
		c.Cache[key] = element
	}
	return nil
}

func (c *inMemory) Get(key string) (any, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if elem, found := c.Cache[key]; found {
		c.DDL.MoveToFront(elem)
		return elem.Value.(*node).Value, nil
	}
	return nil, errors.New(`value not found`)
}

func (c *inMemory) Delete(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if value, found := c.Cache[key]; found {
		c.DDL.Remove(value)
		return nil
	}
	return errors.New(`key not found`)
}

func (c *inMemory) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.Cache = make(map[string]*list.Element)
	c.DDL.Init()
}

func (c *inMemory) ExpiryWorker() {
	for {
		for key, elem := range c.Cache {
			if time.Now().After(elem.Value.(*node).TTL) {
				delete(c.Cache, key)
				c.DDL.Remove(elem)
			}
		}
	}
}

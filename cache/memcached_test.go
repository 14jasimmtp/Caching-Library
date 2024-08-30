package cache

import (
	"fmt"
	"testing"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMemcached_SetGet(t *testing.T) {
	cache := NewMemcachedCache("localhost:11211")

	key := "testKey"
	value := "testValue"
	ttl := time.Second * 10

	err := cache.Set(key, []byte(value), ttl)
	assert.NoError(t, err, "Expected no error when setting a value")

	retrievedValue, err := cache.Get(key)
	assert.NoError(t, err, "Expected no error when getting a value")
	assert.Equal(t, value, retrievedValue, "Expected the retrieved value to match the set value")
}

func TestMemcached_Delete(t *testing.T) {
	cache := NewMemcachedCache("localhost:11211")

	key := "testKeyToDelete"
	value := "testValueToDelete"
	ttl := time.Second * 10

	err := cache.Set(key, []byte(value), ttl)
	assert.NoError(t, err, "Expected no error when setting a value")

	err = cache.Delete(key)
	assert.NoError(t, err, "Expected no error when deleting a value")

	_, err = cache.Get(key)
	assert.Error(t, err, "Expected an error when getting a deleted value")
}

func TestMemcached_Clear(t *testing.T) {
	cache := NewMemcachedCache("localhost:11211")

	key1 := "testKey1"
	value1 := "testValue1"
	key2 := "testKey2"
	value2 := "testValue2"
	ttl := time.Second * 10

	err := cache.Set(key1, []byte(value1), ttl)
	assert.NoError(t, err, "Expected no error when setting the first value")

	err = cache.Set(key2, []byte(value2), ttl)
	assert.NoError(t, err, "Expected no error when setting the second value")

	cache.Clear()

	_, err = cache.Get(key1)
	assert.Error(t, err, "Expected an error when getting a value after clear")

	_, err = cache.Get(key2)
	assert.Error(t, err, "Expected an error when getting a value after clear")
}

func TestGetMiss(t *testing.T) {
	cache := NewMemcachedCache("localhost:11211")

	_, err := cache.Get("nonExistentKey")
	require.Error(t, err)
	assert.Equal(t, memcache.ErrCacheMiss, err)
}

func BenchmarkMemcached_Set(b *testing.B) {
	cache := NewMemcachedCache("localhost:11211")
	key := "benchmarkKey"
	value := []byte("benchmarkValue")
	ttl := time.Second * 10

	for i := 0; i < b.N; i++ {
		err := cache.Set(key, value, ttl)
		if err != nil {
			b.Errorf("Error in Set: %v", err)
		}
	}
}

func BenchmarkMemcached_Get(b *testing.B) {
	cache := NewMemcachedCache("localhost:11211")
	key := "benchmarkKey"
	value := []byte("benchmarkValue")
	ttl := time.Second * 10

	err := cache.Set(key, value, ttl)
	if err != nil {
		b.Fatalf("Error in Set before benchmarking Get: %v", err)
	}

	for i := 0; i < b.N; i++ {
		_, err := cache.Get(key)
		if err != nil {
			b.Errorf("Error in Get: %v", err)
		}
	}
}

func BenchmarkMemcached_Delete(b *testing.B) {
	cache := NewMemcachedCache("localhost:11211")
	key := "benchmarkKey"
	value := []byte("benchmarkValue")
	ttl := time.Second * 10

	err := cache.Set(key, value, ttl)
	if err != nil {
		b.Fatalf("Error in Set before benchmarking Delete: %v", err)
	}
	fmt.Println(b.N)

	for i := 0; i < b.N; i++ {
		err := cache.Delete(key)
		if err != nil {
			b.Errorf("Error in Delete: %v", err)
		}
		err = cache.Set(key, value, ttl)
		if err != nil {
			b.Fatalf("Error in Set after Delete: %v", err)
		}
	}
}

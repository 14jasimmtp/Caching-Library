package cache

import (
	"strconv"
	"testing"
	"time"
)

func TestInMemoryIntegration(t *testing.T) {
	InMemory := NewInMemoryCache(10)
	InMemory.Set("key1", "value1", 5*time.Minute)
	val, err := InMemory.Get("key1")
	if err != nil {
		t.Error(err)
	}
	if val != "value1" {
		t.Error("values not matching")
	}
	err = InMemory.Delete("key1")
	if err != nil {
		t.Error("can't delete", err)
	}

}

func BenchmarkInMemorySet(b *testing.B) {
	cache := NewInMemoryCache(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Set("key", "value", time.Minute)
	}
}

func TestSetAndGet(t *testing.T) {
	cach := NewInMemoryCache(2)

	err := cach.Set("key1", "value1", 10*time.Second)
	if err != nil {
		t.Fatalf("Set failed: %v", err)
	}

	value, err := cach.Get("key1")
	if err != nil {
		t.Fatalf("Get failed: %v", err)
	}
	if value != "value1" {
		t.Errorf("Expected value1, got %v", value)
	}
}

func TestEviction(t *testing.T) {
	cach := NewInMemoryCache(2)

	_ = cach.Set("key1", "value1", 10*time.Second)
	_ = cach.Set("key2", "value2", 10*time.Second)

	_ = cach.Set("key3", "value3", 10*time.Second)

	_, err := cach.Get("key1")
	if err == nil {
		t.Fatalf("Expected key1 to be evicted, but it was found")
	}

	value, err := cach.Get("key3")
	if err != nil {
		t.Fatalf("Expected key3 to be in the cache, but got error: %v", err)
	}
	if value != "value3" {
		t.Errorf("Expected value3, got %v", value)
	}
}

func TestTTLExpiration(t *testing.T) {
	cach := NewInMemoryCache(2)

	_ = cach.Set("key1", "value1", 1*time.Second)
	time.Sleep(2 * time.Second)

	_, err := cach.Get("key1")
	if err == nil {
		t.Fatalf("Expected key1 to be expired, but it was found")
	}
}

func BenchmarkInMemory_Get(b *testing.B) {
	cache := NewInMemoryCache(1000)
	key := "benchmarkKey"
	value := "benchmarkValue"
	ttl := time.Second * 10

	err := cache.Set(key, value, ttl)
	if err != nil {
		b.Fatalf("Error in Set before benchmarking Get: %v", err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := cache.Get(key)
		if err != nil {
			b.Errorf("Error in Get: %v", err)
		}
	}
}

func BenchmarkInMemory_Delete(b *testing.B) {
	cache := NewInMemoryCache(100)
	key := "benchmarkKey"
	value := "benchmarkValue"
	ttl := time.Second * 10

	err := cache.Set(key, value, ttl)
	if err != nil {
		b.Fatalf("Error in Set before benchmarking Delete: %v", err)
	}

	b.ResetTimer()

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

func BenchmarkInMemory_Clear(b *testing.B) {
	cache := NewInMemoryCache(100)

	for i := 0; i < 100; i++ {
		key := "key" + strconv.Itoa(i)
		value := "value" + strconv.Itoa(i)
		ttl := time.Second * 10
		cache.Set(key, value, ttl)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		cache.Clear()
	}
}

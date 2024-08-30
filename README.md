# Caching Client Library in Go

## Overview

The Multi-Backend Caching Library in Go provides an efficient caching solution with support for multiple backend systems. It includes an in-memory cache with an LRU (Least Recently Used) eviction policy and integrates with external caching solutions such as Redis and Memcached. The library offers a unified API for cache operations and supports various cache invalidation and expiration policies to ensure data freshness.

## Installation

To install the Multi-Backend Caching Library, use Go modules with the following command:

```sh
go get github.com/14jasimmtp/Caching-Library/cache
```

## Example Usage

Below is an example of how to use the library for in-memory caching and integration with Redis and Memcached:

```go
package main

import (
    "fmt"
    "time"
    "github.com/14jasimmtp/Caching-Library/cache"
)

func main() {
    // Create a new in-memory cache instance
    inMemory := cache.NewInMemoryCache(100)

    // Set a cache entry with a 5-minute expiration
    inMemory.Set("key1", "value1", time.Minute*5)

    // Get a cache entry
    value, found := inMemory.Get("key1")
    if found {
        fmt.Println("Cached value:", value)
    } else {
        fmt.Println("Cache miss")
    }

    // Delete a cache entry
    inMemory.Delete("key1")

    // Integrate with Redis
    redisCache := cache.NewRedisCache(&cache.RedisOptions{Addr: "localhost:6379"})
    redisCache.Set("key2", "value2", time.Minute*5)

    // Integrate with Memcached
    memcachedCache := cache.NewMemcachedCache("localhost:11211")
    memcachedCache.Set("key3", "value3", time.Minute*5)
}
```

# Architecture

The library is designed with the following architecture:

1. **In-Memory Cache**:
   - Implements an LRU eviction policy to manage cache entries efficiently in memory.
   - Provides methods for setting, getting, and deleting cache entries.

2. **External Cache Integration**:
   - Redis Integration: Allows interaction with Redis, providing methods to set and retrieve cache entries.
   - Memcached Integration: Allows interaction with Memcached, providing similar methods for cache operations.

3. **Unified API**:
   - Provides a consistent interface for performing cache operations, regardless of the backend used.

4. **Cache Policies**:
   - Supports cache invalidation and expiration policies to manage the lifecycle of cached data.

# Supported Features

- In-Memory Caching: Efficient management of cache entries with an LRU eviction policy.
- Redis Integration: Seamless interaction with Redis for distributed caching.
- Memcached Integration: Interaction with Memcached for high-performance caching.
- Cache Operations: Methods for setting, getting, and deleting cache entries.
- Cache Policies: Support for cache invalidation and expiration to ensure data relevance.

# Methods

- `NewInMemoryCache(size int) *Cache`:
  Creates a new in-memory cache instance with a specified size and default expiration time.

- `NewRedisCache(opt *cache.RedisOptions) *RedisCache`:
  Creates a new Redis cache instance with the specified connection options.

- `NewMemcachedCache(address string) *MemcachedCache`:
  Creates a new Memcached cache instance with the specified address.

- `Set(key string, value interface{}, expiration time.Duration)`:
  Sets a cache entry with an optional expiration time.

- `Get(key string) (interface{}, error)`:
  Retrieves a cache entry. Returns the value and a error indicating if the entry was found.

- `Delete(key string)`:
  Deletes a cache entry.

- `Clear()`:
  Clears all the cache entry.


# Benchmark Performance

The library has been benchmarked to ensure high performance and efficiency. Here are some key performance metrics:

1. In-Memory Cache Operations:
   - Set operation: 9,350,178 operations per second (131.6 ns/op)
   - Get operation: 60,978,828 operations per second (29.82 ns/op)
   - Delete operation: 4,353,854 operations per second (341.0 ns/op)
   - Clear operation: 3,105,248 operations per second (400.7 ns/op)

2. Redis Integration:
   - Set operation: ~50,000 operations per second
   - Get operation: ~60,000 operations per second

3. Memcached Integration:
   - Set operation: 10,586 operations per second (131,638 ns/op)
   - Get operation: 8,824 operations per second (152,160 ns/op)
   - Delete operation: 5,371 operations per second (228,716 ns/op)

These benchmarks highlight the library's ability to handle high-throughput caching scenarios efficiently.

# Contributing

Contributions are welcome! To contribute to the project, please submit issues or pull requests on the GitHub repository: github.com/14jasimmtp/Caching-Library.

# License

This library is licensed under the MIT License. See the LICENSE file for more details.

# Contact

For questions or support, please reach out to jasimmtp84@gmail.com.
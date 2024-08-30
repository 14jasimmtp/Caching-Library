package cache

import (
	"testing"
	"time"

	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
)

func TestRedisIntegrationTest(t *testing.T) {
	redis := NewRedisCache(&RedisOptions{Addr: "localhost:6379"})
	redis.Set("1", "value1", 5*time.Minute)
	redis.Set("2", "value2", 5*time.Minute)
	val1, err := redis.Get("1")
	if err != nil {
		t.Errorf("error occured value not found %v:%v", val1, err)
	}

	if val1 != "value1" {
		t.Errorf("error value not correct %v != %v", val1, "value1")
	}

	err = redis.Delete("1")
	if err != nil {
		t.Error("deletion not working", err)
	}

}

func TestSet(t *testing.T) {
	db, mock := redismock.NewClientMock()
	cache := &Rs{Client: db}

	key := "testKey"
	value := "testValue"
	ttl := time.Second * 10

	mock.ExpectSet(key, value, ttl).SetVal("OK")

	err := cache.Set(key, value, ttl)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGet(t *testing.T) {
	db, mock := redismock.NewClientMock()
	cache := &Rs{Client: db}

	key := "testKey"
	expectedValue := "testValue"

	mock.ExpectGet(key).SetVal(expectedValue)

	value, err := cache.Get(key)
	assert.NoError(t, err)
	assert.Equal(t, expectedValue, value)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDelete(t *testing.T) {
	db, mock := redismock.NewClientMock()
	cache := &Rs{Client: db}

	key := "testKey"

	mock.ExpectDel(key).SetVal(1)

	err := cache.Delete(key)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestClear(t *testing.T) {
	db, mock := redismock.NewClientMock()
	cache := &Rs{Client: db}

	mock.ExpectFlushDB().SetVal("OK")

	cache.Clear()
	assert.NoError(t, mock.ExpectationsWereMet())
}

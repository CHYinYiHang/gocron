package cache

import (
	"github.com/coocood/freecache"
)

var cache = freecache.NewCache(1024 * 1024)

func Set(key, value []byte, expireSeconds int) error {
	return cache.Set(key, value, expireSeconds)
}

func Get(key []byte) ([]byte, error) {
	return cache.Get(key)
}

func GetAll(key []byte) ([]byte, uint32, error) {
	return cache.GetWithExpiration(key)
}

func Del(key []byte) bool {
	return cache.Del(key)
}

package cache

import (
	"github.com/bluele/gcache"
)

type Cache struct {
	gcache.Cache
}

func NewCache() *Cache {
	return &Cache{Cache: gcache.New(20000).LRU().Build()}
}

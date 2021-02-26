package cache

import (
	"testing"
)

func TestGCache(t *testing.T) {
	cache := NewCache()
	_ = cache.Set("ok", "ok")
	get, err := cache.Get("ok")
	t.Log(get, err)
}

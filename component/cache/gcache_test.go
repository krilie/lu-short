package cache

import (
	"github.com/bluele/gcache"
	"testing"
)

func TestGCache(t *testing.T) {
	cache := NewCache()
	_ = cache.Set("ok", "ok")
	get, err := cache.Get("ok")
	t.Log(get, err)

	// with cache build
	build := gcache.New(2000).LRU().LoaderFunc(func(key interface{}) (interface{}, error) {
		return key, nil
	}).Build()
	i, err := build.Get("3342")
	t.Log(i, err)
	i2, err := build.Get(234)
	t.Log(i2, err)
	remove := build.Remove(234)
	t.Log(remove)
}

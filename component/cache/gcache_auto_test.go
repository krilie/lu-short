package cache

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestAutoGCache(t *testing.T) {
	cache := NewCache()
	err := cache.Set("ok", "ok")
	assert.Equal(t, nil, err, "should no error", err)
	get, err := cache.Get("ok")
	assert.Equal(t, nil, err, "should no error", err)
	assert.Equal(t, "ok", get, "should a ok str", errors.New("this is err"))
	fmt.Println(get, err)
}

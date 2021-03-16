package nredis

import (
	"lu-short/common/appdig"
	"lu-short/component"
	"testing"
	"time"
)

func TestNewNRedis(t *testing.T) {
	dig := appdig.NewAppDig().
		MustProvides(component.DigComponentProviderAll).
		MustProvide(NewNRedis)

	dig.MustInvoke(func(rd *NRedis) {
		rd.client.Ping()
		rd.client.Set("aaa", "bbb", time.Second*89)

	})
}

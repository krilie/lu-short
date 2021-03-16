package ndb

import (
	"lu-short/common/appdig"
	"lu-short/component/ncfg"
	"lu-short/component/nlog"
	"testing"
)

//go:generate go test -v ./...

func TestNewNDb(t *testing.T) {
	dig := appdig.NewAppDig()
	dig.MustProvide(ncfg.NewNConfig)
	dig.MustProvide(NewNDb)
	dig.MustProvide(nlog.NewNLog)

	dig.MustInvoke(func(ndb *NDb) {
		ndb.Ping()
	})
}

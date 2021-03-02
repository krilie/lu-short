package dao

import (
	"lu-short/common/appdig"
	"lu-short/component"
	"testing"
)

func TestAutoNewLuShortDao(t *testing.T) {
	dig := appdig.NewAppDig()
	dig.MustProvides(component.DigComponentProviderAll)
	dig.MustProvide(NewLuShortDao)

	dig.MustInvoke(func(dao *LuShortDao) {
		dao.dao.Ping()
	})
}

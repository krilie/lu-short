package dao

import (
	"lu-short/component/ndb"
	"lu-short/component/nlog"
)

type LuShortDao struct {
	dao *ndb.NDb
	log *nlog.NLog
}

func NewLuShortDao(dao *ndb.NDb, log *nlog.NLog) *LuShortDao {
	return &LuShortDao{
		dao: dao,
		log: log,
	}
}

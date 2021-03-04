package visitor_info_dao

import (
	"lu-short/component/ndb"
	"lu-short/component/nlog"
)

type VisitorInfoDao struct {
	dao *ndb.NDb
	log *nlog.NLog
}

func NewVisitorInfoDao(dao *ndb.NDb, log *nlog.NLog) *VisitorInfoDao {
	return &VisitorInfoDao{
		dao: dao,
		log: log,
	}
}

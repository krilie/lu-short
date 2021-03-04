package service

import (
	"lu-short/component/nlog"
	"lu-short/module/dao/lushort_dao"
)

type LuShortService struct {
	Dao *lushort_dao.LuShortDao
	Log *nlog.NLog
}

func NewLuShortService(dao *lushort_dao.LuShortDao, log *nlog.NLog) *LuShortService {
	return &LuShortService{
		Dao: dao,
		Log: log,
	}
}

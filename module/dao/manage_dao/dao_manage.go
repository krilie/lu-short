package manage_dao

import (
	"context"
	"lu-short/component/ndb"
	"lu-short/component/nlog"
	"lu-short/module/model"
)

type ManageDao struct {
	dao *ndb.NDb
	log *nlog.NLog
}

func NewManageDao(dao *ndb.NDb, log *nlog.NLog) *ManageDao {
	return &ManageDao{
		dao: dao,
		log: log,
	}
}

func (dao *ManageDao) CreateManage(ctx context.Context, manage *model.TbManage) error {
	return nil
}

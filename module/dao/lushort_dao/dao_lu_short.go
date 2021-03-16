package lushort_dao

import (
	"context"
	_ "embed"
	"lu-short/common/utils/id_util"
	"lu-short/component/ndb"
	"lu-short/component/nlog"
	"lu-short/module/model"
	"time"
)

type LuShortDao struct {
	dao *ndb.NDb
	log *nlog.NLog
}

func NewLuShortDao(dao *ndb.NDb, log *nlog.NLog) *LuShortDao {
	var luShortDao = &LuShortDao{
		dao: dao,
		log: log,
	}

	// 设置表名
	luShortDao.dao.AddTable(model.TbRedirect{}, "tb_redirect").SetKeys(false, "id")
	luShortDao.dao.AddTable(model.TbRedirectLog{}, "tb_redirect_log").SetKeys(false, "id")

	return luShortDao
}

func (dao *LuShortDao) GetReDirectById(ctx context.Context, id interface{}) (m *model.TbRedirect, err error) {
	m = &model.TbRedirect{}
	err = dao.dao.Get(ctx, m, "select * from tb_redirect where deleted_at is null and `id`=?", id)
	return m, err
}

func (dao *LuShortDao) GetReDirectByKey(ctx context.Context, key interface{}) (m *model.TbRedirect, err error) {
	m = &model.TbRedirect{}
	err = dao.dao.GetDb(ctx).SelectOne(m, "select * from tb_redirect where deleted_at is null and `key`=?", key)
	return m, err
}

func (dao *LuShortDao) UpdateReDirect(ctx context.Context, m *model.TbRedirect) error {

	affected, err := dao.dao.GetDb(ctx).Update(m)
	if err != nil && affected > 0 {
		return err
	}

	return nil
}

func (dao *LuShortDao) DeleteLuShort(ctx context.Context, id, key string) error {
	_, err := dao.dao.Exec(ctx, "update tb_redirect set deleted_at=? where id=? and `key`=?", time.Now(), id, key)
	return err
}

func (dao *LuShortDao) CreateLuShort(ctx context.Context, m *model.TbRedirect) error {
	if m.Id == "" {
		m.Id = id_util.GetUuid()
	}
	return dao.dao.GetDb(ctx).Insert(m)
}

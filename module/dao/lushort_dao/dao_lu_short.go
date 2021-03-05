package lushort_dao

import (
	"context"
	_ "embed"
	"github.com/bluele/gcache"
	"lu-short/common/utils/id_util"
	"lu-short/component/ndb"
	"lu-short/component/nlog"
	"lu-short/module/model"
	"time"
)

type LuShortDao struct {
	dao              *ndb.NDb
	log              *nlog.NLog
	directCatchByKey gcache.Cache
}

func NewLuShortDao(dao *ndb.NDb, log *nlog.NLog) *LuShortDao {
	var luShortDao = &LuShortDao{
		dao:              dao,
		log:              log,
		directCatchByKey: nil,
	}

	luShortDao.directCatchByKey = gcache.New(200).
		LoaderFunc(func(key interface{}) (interface{}, error) {
			return luShortDao.getReDirectByKey(context.Background(), key)
		}).
		Build()

	// 设置表名
	luShortDao.dao.AddTable(model.TbRedirect{}, "tb_redirect").SetKeys(false, "id")
	luShortDao.dao.AddTable(model.TbRedirectLog{}, "tb_redirect_log").SetKeys(false, "id")

	return luShortDao
}

func (dao *LuShortDao) getReDirectById(ctx context.Context, id interface{}) (m *model.TbRedirect, err error) {
	m = &model.TbRedirect{}
	err = dao.dao.Get(ctx, m, "select * from tb_redirect where deleted_at is null and `id`=?", id)
	return m, err
}

func (dao *LuShortDao) getReDirectByKey(ctx context.Context, key interface{}) (m *model.TbRedirect, err error) {
	m = &model.TbRedirect{}
	err = dao.dao.GetDb(ctx).SelectOne(m, "select * from tb_redirect where deleted_at is null and `key`=?", key)
	return m, err
}

func (dao *LuShortDao) GetReDirectById(ctx context.Context, id interface{}) (m *model.TbRedirect, err error) {
	return dao.getReDirectById(ctx, id)
}

func (dao *LuShortDao) GetReDirectByKey(ctx context.Context, key interface{}) (m *model.TbRedirect, err error) {
	m = &model.TbRedirect{}
	redirect, err := dao.directCatchByKey.Get(key)
	if err != nil {
		return nil, err
	}
	return redirect.(*model.TbRedirect), nil
}

func (dao *LuShortDao) UpdateReDirect(ctx context.Context, m *model.TbRedirect) error {

	affected, err := dao.dao.GetDb(ctx).Update(m)
	if err != nil && affected > 0 {
		return err
	}

	dao.directCatchByKey.Remove(m.Key)

	return nil
}

func (dao *LuShortDao) DeleteLuShort(ctx context.Context, id, key string) error {
	_, err := dao.dao.Exec(ctx, "update tb_redirect set deleted_at=? where id=? and `key`=?", time.Now(), id, key)
	if err == nil {
		dao.directCatchByKey.Remove(key)
	}
	return err
}

func (dao *LuShortDao) CreateLuShort(ctx context.Context, m *model.TbRedirect) error {
	if m.Id == "" {
		m.Id = id_util.GetUuid()
	}
	return dao.dao.GetDb(ctx).Insert(m)
}

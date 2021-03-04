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
	dao         *ndb.NDb
	log         *nlog.NLog
	directCatch gcache.Cache
}

func NewLuShortDao(dao *ndb.NDb, log *nlog.NLog) *LuShortDao {
	var luShortDao = &LuShortDao{
		dao:         dao,
		log:         log,
		directCatch: nil,
	}

	luShortDao.directCatch = gcache.New(200).
		LoaderFunc(func(key interface{}) (interface{}, error) {
			return luShortDao.getReDirectById(context.Background(), key)
		}).
		Build()

	// 设置表名
	luShortDao.dao.AddTable(model.TbRedirect{}, "tb_redirect").SetKeys(false, "id")
	luShortDao.dao.AddTable(model.TbRedirectLog{}, "tb_redirect_log").SetKeys(false, "id")

	return luShortDao
}

func (dao *LuShortDao) getReDirectById(ctx context.Context, id interface{}) (model *model.TbRedirect, err error) {
	err = dao.dao.Get(ctx, model, "", id)
	return model, err
}

func (dao *LuShortDao) GetReDirectById(ctx context.Context, id interface{}) (m *model.TbRedirect, err error) {
	redirect, err := dao.directCatch.Get(id)
	return redirect.(*model.TbRedirect), err
}

func (dao *LuShortDao) UpdateReDirect(ctx context.Context, m *model.TbRedirect) error {

	affected, err := dao.dao.GetDb(ctx).Update(m)
	if err != nil && affected > 0 {
		return err
	}

	dao.directCatch.Remove(m.Id)

	return nil
}

func (dao *LuShortDao) DeleteLuShort(ctx context.Context, id string) error {
	_, err := dao.dao.Exec(ctx, "update table tb_redirect set deleted_at=? where id=?", time.Now(), id)
	if err == nil {
		dao.directCatch.Remove(id)
	}
	return err
}

func (dao *LuShortDao) CreateLuShort(ctx context.Context, m *model.TbRedirect) error {
	if m.Id == "" {
		m.Id = id_util.GetUuid()
	}
	return dao.dao.GetDb(ctx).Insert(m)
}

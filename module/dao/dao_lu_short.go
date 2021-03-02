package dao

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/bluele/gcache"
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

	sql, args, err := sq.Update("tb_redirect").
		Where("id=?", m.Id).Where("deleted_at is null").
		Set("updated_at=?", time.Now()).
		Set("customer_id=?", m.Id).
		Set("ori_url=?", m.Id).
		Set("key=?", m.Id).
		Set("rate_limit=?", m.Id).
		Set("times_limit_left=?", m.Id).
		Set("jump_limit_left=?", m.Id).
		Set("begin_time=?", m.Id).
		Set("dead_time=?", m.Id).
		ToSql()
	if err != nil {
		panic(err)
	}

	affected, err := dao.dao.Exec(ctx, sql, args...)
	if err != nil && affected > 0 {
		return err
	}

	dao.directCatch.Remove(m.Id)

	return nil
}

func (dao *LuShortDao) DeleteLuShort(ctx context.Context, id string) error {
	_, err := dao.dao.Exec(ctx, "update table tb_redirect set deleted_at=? where id=?", time.Now(), id)
	return err
}

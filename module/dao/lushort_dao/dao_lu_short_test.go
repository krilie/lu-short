package lushort_dao

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"lu-short/common/appdig"
	"lu-short/common/com_model"
	"lu-short/common/utils/id_util"
	"lu-short/common/utils/random"
	"lu-short/component"
	"lu-short/module/model"
	"testing"
	"time"
)

func TestAutoNewLuShortDao(t *testing.T) {
	dig := appdig.NewAppDig()
	dig.MustProvides(component.DigComponentProviderAll)
	dig.MustProvide(NewLuShortDao)

	dig.MustInvoke(func(dao *LuShortDao) {
		redirect := RandomTbRedirect()
		err := dao.CreateLuShort(context.Background(), redirect)
		require.Nil(t, err)
		redirect.Key = "123456"
		err = dao.UpdateReDirect(context.Background(), redirect)
		require.Nil(t, err)
		get, err := dao.dao.GetDb(context.Background()).Get(&model.TbRedirect{}, redirect.Id)
		require.Nil(t, err)
		assert.Equal(t, redirect.Key, get.(*model.TbRedirect).Key)
	})
}

func RandomTbRedirect() *model.TbRedirect {
	return &model.TbRedirect{
		TbCommon: com_model.TbCommon{
			Id:        id_util.GetUuid(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
		},
		CustomerId:     random.GetRandomStr(6),
		OriUrl:         random.GetRandomStr(6),
		Key:            random.GetRandomStr(6),
		RateLimit:      random.GetRandomInt32(),
		TimesLimitLeft: random.GetRandomInt32(),
		JumpLimitLeft:  random.GetRandomInt32(),
		BeginTime:      time.Now(),
		EndTime:        time.Now(),
	}
}

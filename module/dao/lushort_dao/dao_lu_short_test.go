package lushort_dao

import (
	"context"
	"github.com/stretchr/testify/require"
	"lu-short/common/appdig"
	"lu-short/common/com_model"
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
		err := dao.CreateLuShort(context.Background(), &model.TbRedirect{
			TbCommon: com_model.TbCommon{
				Id:        "",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				DeletedAt: nil,
			},
			CustomerId:     "1",
			OriUrl:         "2",
			Key:            "3",
			RateLimit:      4,
			TimesLimitLeft: 5,
			JumpLimitLeft:  6,
			BeginTime:      time.Now(),
			EndTime:        time.Now(),
		})
		require.Nil(t, err)
	})
}

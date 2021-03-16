package service

import (
	"context"
	"errors"
	"github.com/bluele/gcache"
	"github.com/juju/ratelimit"
	"lu-short/common/errs"
	"lu-short/component/nlog"
	"lu-short/module/dao/lushort_dao"
	"sync/atomic"
	"time"
)

type LuShortService struct {
	Dao   *lushort_dao.LuShortDao
	Log   *nlog.NLog
	Limit gcache.Cache
}

func NewLuShortService(dao *lushort_dao.LuShortDao, log *nlog.NLog) *LuShortService {
	return &LuShortService{
		Dao:   dao,
		Log:   log,
		Limit: gcache.New(2000).LRU().Build(),
	}
}

type LimitInfo struct {
	Id        string
	Key       string
	OriUrl    string
	RateLimit *ratelimit.Bucket
	leftJump  int64
}

func (l *LimitInfo) SetLeftJump(jump int64) {
	atomic.StoreInt64(&l.leftJump, jump)
}
func (l *LimitInfo) GetLeftJump() int64 {
	loadInt64 := atomic.LoadInt64(&l.leftJump)
	return loadInt64
}
func (l *LimitInfo) AddLeftJump(add int64) int64 {
	return atomic.AddInt64(&l.leftJump, add)
}

func (s *LuShortService) Redirect(ctx context.Context, key, customerId, ip, agent, deviceId string) (oriUrl string, err error) {
	// 短域名 =>限速 限次 限个数 .. 限ip(段) 禁id(多个) 禁ip(段)
	var loadLimitInfo = func(key string) *LimitInfo {
		load, err := s.Limit.Get(key)
		if err != nil && !errors.Is(err, gcache.KeyNotFoundError) {
			panic(err)
		}
		if load == nil {
			redirect, err := s.Dao.GetReDirectByKey(ctx, key)
			if err != nil {
				panic(err)
			}
			if redirect == nil {
				return nil
			}
			var limitInfo = &LimitInfo{
				Id:        redirect.Id,
				Key:       redirect.Key,
				OriUrl:    redirect.OriUrl,
				RateLimit: ratelimit.NewBucketWithRate(float64(redirect.RateLimit), 10),
				leftJump:  redirect.TimesLimitLeft,
			}
			err = s.Limit.Set(key, limitInfo)
			if err != nil {
				panic(err)
			}
			return limitInfo
		} else {
			return load.(*LimitInfo)
		}
	}
	// 获取
	limitInfo := loadLimitInfo(key)
	if limitInfo == nil {
		return "", errs.NewNotExistsError().WithMsg("not find")
	}
	// 限速
	_, b := limitInfo.RateLimit.TakeMaxDuration(1, time.Second*3)
	if !b {
		return "", errs.NewNormal().WithMsg("temp err, try again place")
	}
	// 限次
	jump := limitInfo.AddLeftJump(-1)
	if jump <= 0 {
		return "", errs.NewNormal().WithMsg("temp err, no times left")
	}
	// todo: 消息队列 次数减一操作 param: reduce-1 customerId ip
	return limitInfo.OriUrl, nil
}

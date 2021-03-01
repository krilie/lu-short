package component

import (
	"lu-short/component/broker"
	"lu-short/component/cache"
	"lu-short/component/cron"
	"lu-short/component/ncfg"
	"lu-short/component/ndb"
	"lu-short/component/nlog"
)

// 生产环境 开发环境&测试环境 通过配置文件区分
var DigComponentProviderAll = []interface{}{
	ncfg.NewNConfig,
	nlog.NewNLog,
	ndb.NewNDb,
	broker.NewBroker,
	cache.NewCache,
	cron.NewCrone,
}

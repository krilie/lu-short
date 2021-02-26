package component

import (
	"lu-short/component/broker"
	"lu-short/component/cache"
	"lu-short/component/cron"
	"lu-short/component/ncfg"
	"lu-short/component/ndb"
	"lu-short/component/nlog"
)

var DigComponentProviderAll = []interface{}{
	ncfg.NewNConfigByFileFromEnv("APP_CONFIG_PATH"),
	nlog.NewLogger,
	ndb.NewNDb,
	broker.NewBroker,
	cache.NewCache,
	cron.NewCrone,
}

var DigComponentProviderAllForTest = []interface{}{
	ncfg.NewNConfigByCfgStrFromEnvJson("MYAPP_TEST_CONFIG"),
	nlog.NewLogger,
	ndb.NewNDb,
	broker.NewBroker,
	cache.NewCache,
	cron.NewCrone,
}

package nredis

import (
	"github.com/go-redis/redis/v7"
	"lu-short/component/ncfg"
)

type NRedis struct {
	client *redis.Client
}

func NewNRedis(cfg *ncfg.NConfig) *NRedis {
	var redisCfg = cfg.GetRedisCfg()

	var retRedis = &NRedis{client: redis.NewClient(&redis.Options{
		Addr:     redisCfg.ConnStr,
		Password: redisCfg.Auth, // no password set
		DB:       0,             // use default DB
	})}
	_, err := retRedis.client.Ping().Result()
	if err != nil {
		panic(err)
	}
	return retRedis
}

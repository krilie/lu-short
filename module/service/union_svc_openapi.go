package service

import "lu-short/component/nlog"

type UnionSvcOpenApi struct {
	LuShortService *LuShortService
	log            *nlog.NLog
}

func NewUnionSvcOpenApi(luShortService *LuShortService, log *nlog.NLog) *UnionSvcOpenApi {
	return &UnionSvcOpenApi{LuShortService: luShortService, log: log}
}

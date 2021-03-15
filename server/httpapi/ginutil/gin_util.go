package ginutil

import (
	context2 "context"
	"github.com/gin-gonic/gin"
	"lu-short/component/nlog"
)

// some const value for gin http protocol
var (
	GinKeyAppContext = "GinKeyAppContext"
)

type GinWrap struct {
	log *nlog.NLog
	*gin.Context
	AppCtx context2.Context
}

func NewGinWrap(ctx *gin.Context, log *nlog.NLog) *GinWrap {
	var wrap = &GinWrap{log: log, Context: ctx}
	wrap.AppCtx = wrap.GetAppContext()
	return wrap
}

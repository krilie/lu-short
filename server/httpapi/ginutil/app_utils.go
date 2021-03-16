package ginutil

import (
	context2 "context"
	"golang.org/x/net/context"
	"lu-short/common/utils/id_util"
)

func (g *GinWrap) GetAppContext() context2.Context {
	ctx := context.WithValue(context.Background(), "ginCtx", g.Context)
	ctx = context.WithValue(ctx, "trace_id", id_util.GetUuid())
	ctx = context.WithValue(ctx, "ip", g.ClientIP())
	return ctx
}

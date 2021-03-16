package main

import (
	"context"
	"lu-short/common/appdig"
	"lu-short/common/utils/sigutil"
	"lu-short/component"
	"lu-short/component/ndb"
	"lu-short/component/nlog"
	"lu-short/module/dao"
	"lu-short/module/service"
	"lu-short/server"
	"lu-short/server/httpapi"
	"time"
)

// 短域名 =>限速 限次 限个数 .. 限ip(段) 禁id(多个) 禁ip(段)

//go:generate swag init -g ./main.go  # install: go get -u github.com/swaggo/swag/cmd/swag
//go:generate go test -run Auto -v ./...

func main() {
	dig := appdig.NewAppDig()
	dig.MustProvides(component.DigComponentProviderAll)
	dig.MustProvides(dao.DaoAll)
	dig.MustProvides(service.SvcAll)
	dig.MustProvide(httpapi.NewHttpOpenApi)
	dig.MustProvide(httpapi.NewHttpApi)
	dig.MustProvide(server.NewService)

	dig.MustInvoke(func(svc *server.Service, log *nlog.NLog, dao *ndb.NDb) {
		var ctx = context.Background()
		defer dao.CloseDb()
		defer log.Close()
		closeSvc := svc.StartService(ctx)
		sigutil.WaitSignalAndExit(ctx, func() {
			err := closeSvc(time.Second * 10)
			if err != nil {
				log.Get(ctx).Sugar().Errorf("err exit %v", err.Error())
			} else {
				log.Get(ctx).Info("exit no err")
			}
		})
	})

}

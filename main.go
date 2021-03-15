package main

import (
	"context"
	"lu-short/common/appdig"
	"lu-short/component"
	"lu-short/component/nlog"
	"lu-short/module/dao"
	"lu-short/module/service"
	"lu-short/server"
	"lu-short/server/httpapi"
	"os"
	"os/signal"
	"syscall"
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

	dig.MustInvoke(func(svc *server.Service, log *nlog.NLog) {
		var ctx = context.Background()
		closeSvc := svc.StartService(ctx)
		WaitSignalAndExit(ctx, func() {
			err := closeSvc(time.Second * 10)
			if err != nil {
				log.Get(ctx).Sugar().Errorf("err exit %v", err.Error())
			} else {
				log.Get(ctx).Info("exit no err")
			}
		})
	})

}

// 接收信号和退出
func WaitSignalAndExit(ctx context.Context, exit func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL:
			exit()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

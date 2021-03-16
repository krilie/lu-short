package sigutil

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

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

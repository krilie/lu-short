package cron

import (
	context2 "context"
	"fmt"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	m.Run()
}

func TestCron(t *testing.T) {
	c := NewCrone()
	_, _ = c.AddFunc("*/1 * * * * *", func() {
		fmt.Println("ok")
	})
	_, _ = c.AddFunc("@every 2s", func() {
		fmt.Println("ok 2")
	})
	time.Sleep(time.Second * 10)
	c.StopAndWait(context2.Background())
}

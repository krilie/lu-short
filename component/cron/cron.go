package cron

import (
	"context"
	"github.com/robfig/cron/v3"
)

type NCron struct {
	*cron.Cron
}

func (c *NCron) StopAndWait(ctx context.Context) {
	if c.Cron != nil {
		stop := c.Cron.Stop()
		<-stop.Done()
	}
}

func NewCrone() *NCron {
	CronGlobal := cron.New(cron.WithParser(cron.NewParser(cron.Second|cron.Minute|cron.Hour|cron.Dom|cron.Month|cron.DowOptional|cron.Descriptor)),
		cron.WithChain(cron.Recover(cron.DefaultLogger)))
	CronGlobal.Start()
	return &NCron{Cron: CronGlobal}
}

func (c *NCron) MustAddFunc(ctx context.Context, spec string, f func()) {
	_, err := c.Cron.AddFunc(spec, f)
	if err != nil {
		panic(err)
	}
}

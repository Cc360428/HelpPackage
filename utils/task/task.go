package task

import (
	"github.com/Cc360428/HelpPackage/utils/logs"
	"github.com/Cc360428/HelpPackage/utils/timec"
	"github.com/robfig/cron/v3"
	"time"
)

func Task() {
	c := cron.New(cron.WithSeconds())
	_, _ = c.AddFunc("0/1 * * * * ?", func() {
		logs.Info(timec.DateDayFormat())
	})
	c.Start()
	time.Sleep(time.Second * 3)
}

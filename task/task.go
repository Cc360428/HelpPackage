package task

import (
	"log"
	"time"

	"github.com/Cc360428/HelpPackage/timec"
	"github.com/robfig/cron/v3"
)

func Task() {
	c := cron.New(cron.WithSeconds())
	_, _ = c.AddFunc("0/1 * * * * ?", func() {
		log.Println(timec.DateDayFormat())
	})
	c.Start()
	time.Sleep(time.Second * 3)
}

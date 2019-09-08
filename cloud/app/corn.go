package app

import (
	"fmt"
	"github.com/robfig/cron"
)

var Cron = new(CronService)

type CronService struct {
	cron *cron.Cron
}

func (s *CronService) StartCorn() (err error) {
	s.cron = cron.New()
	err = s.cron.AddFunc("0 * * * * *", func() {
		fmt.Println("First scheduled task.")
	})
	if err != nil {
		return err
	}

	s.cron.Start()
	return nil
}

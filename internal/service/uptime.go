package service

import (
	"ssm/internal/domain/uptime"
	"time"
)

type UptimeService struct {
	repo   uptime.UptimeRepo
	period int
}

func NewUptimeService(repo uptime.UptimeRepo, period int) *UptimeService {
	return &UptimeService{repo: repo, period: period}
}

func (us *UptimeService) StreamUptime(ch chan *uptime.Uptime) {
	for {
		uptimeInfo, _ := us.repo.GetUptime()
		ch <- uptimeInfo

		time.Sleep(time.Duration(us.period) * time.Second)
	}
}

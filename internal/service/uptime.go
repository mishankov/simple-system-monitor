package service

import (
	"context"
	"time"

	"github.com/mishankov/simple-system-monitor/internal/domain/uptime"
)

type UptimeService struct {
	repo   uptime.Repo
	period int
}

func NewUptimeService(repo uptime.Repo, period int) *UptimeService {
	return &UptimeService{repo: repo, period: period}
}

func (us *UptimeService) StreamUptime(ctx context.Context, ch chan *uptime.Uptime) {
	for {
		uptimeInfo, _ := us.repo.GetUptime()
		ch <- uptimeInfo

		done := false
		select {
		case <-time.After(time.Duration(us.period) * time.Second):
		case <-ctx.Done():
			done = true
		}

		if done {
			close(ch)
			return
		}
	}
}

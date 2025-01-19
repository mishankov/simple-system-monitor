package service

import (
	"context"
	"time"

	"github.com/mishankov/simple-system-monitor/internal/domain/meminfo"
)

type MemInfoService struct {
	repo   meminfo.MemInfoRepo
	period int
}

func NewMemInfoService(repo meminfo.MemInfoRepo, period int) *MemInfoService {
	return &MemInfoService{repo: repo, period: period}
}

func (mis *MemInfoService) StreamMemInfo(ctx context.Context, ch chan *meminfo.MemInfo) {
	for {
		memInfo, _ := mis.repo.GetMemInfo()
		ch <- memInfo

		done := false
		select {
		case <-time.After(time.Duration(mis.period) * time.Second):
		case <-ctx.Done():
			done = true
		}

		if done {
			close(ch)
			return
		}
	}
}

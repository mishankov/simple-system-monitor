package service

import (
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

func (mis *MemInfoService) StreamMemInfo(ch chan *meminfo.MemInfo) {
	for {
		memInfo, _ := mis.repo.GetMemInfo()
		ch <- memInfo

		time.Sleep(time.Duration(mis.period) * time.Second)
	}
}

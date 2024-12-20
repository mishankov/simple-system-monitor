package service

import (
	"ssm/internal/domain/meminfo"
	"time"
)

type MemInfoService struct {
	repo meminfo.MemInfoRepo
}

func NewMemInfoService(repo meminfo.MemInfoRepo) *MemInfoService {
	return &MemInfoService{repo: repo}
}

func (mis *MemInfoService) StreamMemInfo(ch chan *meminfo.MemInfo, period time.Duration) {
	for {
		memInfo, _ := mis.repo.GetMemInfo()
		ch <- memInfo

		time.Sleep(period)
	}
}

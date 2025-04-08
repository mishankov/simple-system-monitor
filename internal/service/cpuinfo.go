package service

import (
	"context"
	"time"

	"github.com/mishankov/simple-system-monitor/internal/domain/cpuinfo"
)

type CPUInfoService struct {
	repo   cpuinfo.Repo
	period int
}

func NewCPUInfoService(repo cpuinfo.Repo, period int) *CPUInfoService {
	return &CPUInfoService{repo: repo, period: period}
}

func (cis *CPUInfoService) StreamCPULoad(ctx context.Context, ch chan []cpuinfo.CPULoad) {
	initial := []cpuinfo.CPULoad{}
	prevData, _ := cis.repo.GetCPUInfo()
	for _, cpuInfo := range prevData {
		initial = append(initial, cpuinfo.CPULoad{ID: cpuInfo.ID, Load: 0})
	}
	ch <- initial

	time.Sleep(time.Duration(cis.period) * time.Second)
	for {
		cpuInfos, _ := cis.repo.GetCPUInfo()

		loads := []cpuinfo.CPULoad{}
		for i, cpuInfo := range cpuInfos {
			// https://stackoverflow.com/a/23376195
			prevIdle := prevData[i].Idle + prevData[i].Iowait
			idle := cpuInfo.Idle + cpuInfo.Iowait

			prevNonIdle := prevData[i].User + prevData[i].Nice + prevData[i].System + prevData[i].Irq + prevData[i].Softirq + prevData[i].Steal
			nonIdle := cpuInfo.User + cpuInfo.Nice + cpuInfo.System + cpuInfo.Irq + cpuInfo.Softirq + cpuInfo.Steal

			prevTotal := prevIdle + prevNonIdle
			total := idle + nonIdle

			totalDiff := total - prevTotal
			idleDiff := idle - prevIdle

			loads = append(loads, cpuinfo.CPULoad{ID: cpuInfo.ID, Load: float32(totalDiff-idleDiff) / float32(totalDiff)})
		}

		ch <- loads

		prevData = cpuInfos

		done := false
		select {
		case <-time.After(time.Duration(cis.period) * time.Second):
		case <-ctx.Done():
			done = true
		}

		if done {
			close(ch)
			return
		}
	}
}

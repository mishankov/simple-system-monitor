package service

import (
	"ssm/internal/domain/cpuinfo"
	"time"
)

type CPUInfoService struct {
	repo cpuinfo.CPUInfoRepo
}

func NewCPUInfoService(repo cpuinfo.CPUInfoRepo) *CPUInfoService {
	return &CPUInfoService{repo: repo}
}

func (cis *CPUInfoService) StreamCPULoad(ch chan []cpuinfo.CPULoad, period time.Duration) {
	initial := []cpuinfo.CPULoad{}
	prevData, _ := cis.repo.GetCPUInfo()
	for _, cpuInfo := range prevData {
		initial = append(initial, cpuinfo.CPULoad{Id: cpuInfo.Id, Load: 0})
	}
	ch <- initial

	time.Sleep(period)
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

			loads = append(loads, cpuinfo.CPULoad{Id: cpuInfo.Id, Load: float32(totalDiff-idleDiff) / float32(totalDiff)})
		}

		ch <- loads

		prevData = cpuInfos
		time.Sleep(period)
	}
}

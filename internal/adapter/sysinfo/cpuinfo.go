package sysinfo

import (
	"bytes"
	"io"
	"os"
	"ssm/internal/domain/cpuinfo"
	"strconv"
	"strings"
)

type CPUInfoRepo struct {
	path string
}

func NewCPUInfoRepo(path string) *CPUInfoRepo {
	return &CPUInfoRepo{path}
}

func (cir *CPUInfoRepo) GetCPUInfo() ([]cpuinfo.CPUInfo, error) {
	file, err := os.Open(cir.path)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	lines := bytes.Split(data, []byte{10})

	cpuInfos := []cpuinfo.CPUInfo{}
	for _, lineByte := range lines {
		line := string(lineByte)

		if !strings.HasPrefix(line, "cpu") || strings.HasPrefix(line, "cpu ") {
			continue
		}

		cpuInfo := cpuinfo.CPUInfo{}

		data := strings.Split(line, " ")

		cpuInfo.Id = data[0]
		cpuInfo.User, _ = strconv.Atoi(data[1])
		cpuInfo.Nice, _ = strconv.Atoi(data[2])
		cpuInfo.System, _ = strconv.Atoi(data[3])
		cpuInfo.Idle, _ = strconv.Atoi(data[4])
		cpuInfo.Iowait, _ = strconv.Atoi(data[5])
		cpuInfo.Irq, _ = strconv.Atoi(data[6])
		cpuInfo.Softirq, _ = strconv.Atoi(data[7])
		cpuInfo.Steal, _ = strconv.Atoi(data[8])
		cpuInfo.Guest, _ = strconv.Atoi(data[9])
		cpuInfo.GuestNice, _ = strconv.Atoi(data[10])

		cpuInfos = append(cpuInfos, cpuInfo)
	}

	return cpuInfos, nil
}

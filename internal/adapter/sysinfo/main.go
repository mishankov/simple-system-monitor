package sysinfo

import (
	"bytes"
	"io"
	"os"
	"ssm/internal/domain/meminfo"
	"strconv"
	"strings"
)

type MemInfoService struct{}

func NewMemInfoService() *MemInfoService {
	return &MemInfoService{}
}

const MEMINFO_PATH = "/proc/meminfo"

func (mis *MemInfoService) GetMemService() (*meminfo.MemInfo, error) {
	file, err := os.Open(MEMINFO_PATH)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	memInfo := &meminfo.MemInfo{}
	for _, lineByte := range bytes.Split(data, []byte{10}) {
		if len(lineByte) == 0 {
			continue
		}

		line := string(lineByte)
		lineSplited := strings.Split(line, ":")
		key := lineSplited[0]
		value, err := strconv.Atoi(strings.Split(strings.TrimSpace(lineSplited[1]), " ")[0])
		if err != nil {
			return nil, err
		}

		switch key {
		case "MemTotal":
			memInfo.MemTotal = value
		case "MemFree":
			memInfo.MemFree = value
		case "MemAvailable":
			memInfo.MemAvailable = value
		}
	}

	return memInfo, nil
}

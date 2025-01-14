package sysinfo

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/mishankov/simple-system-monitor/internal/domain/meminfo"
)

type MemInfoRepo struct {
	dataReader DataReader
}

func NewMemInfoRepo(dataReader DataReader) *MemInfoRepo {
	return &MemInfoRepo{dataReader: dataReader}
}

func (mir *MemInfoRepo) GetMemInfo() (*meminfo.MemInfo, error) {
	data, err := mir.dataReader.ReadData()
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

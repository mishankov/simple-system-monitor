package sysinfo

import (
	"strconv"
	"strings"

	"github.com/mishankov/simple-system-monitor/internal/domain/uptime"
)

type UptimeRepo struct {
	dataReader DataReader
}

func NewUptimeRepo(dataReader DataReader) *UptimeRepo {
	return &UptimeRepo{dataReader: dataReader}
}

func (ur *UptimeRepo) GetUptime() (*uptime.Uptime, error) {
	data, err := ur.dataReader.ReadData()
	if err != nil {
		return nil, err
	}

	uptimeStr := strings.Split(string(data), " ")[0]
	uptimeFloat, err := strconv.ParseFloat(uptimeStr, 32)
	if err != nil {
		return nil, err
	}

	return &uptime.Uptime{Uptime: float32(uptimeFloat)}, nil
}

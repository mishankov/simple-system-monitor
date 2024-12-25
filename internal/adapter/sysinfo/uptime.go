package sysinfo

import (
	"io"
	"os"
	"ssm/internal/domain/uptime"
	"strconv"
	"strings"
)

type UptimeRepo struct {
	path string
}

func NewUptimeRepo(path string) *UptimeRepo {
	return &UptimeRepo{path}
}

func (ur *UptimeRepo) GetUptime() (*uptime.Uptime, error) {
	file, err := os.Open(ur.path)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(file)
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

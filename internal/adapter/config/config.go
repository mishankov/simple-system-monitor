package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getEnvOrDefault(key, def string) string {
	value := os.Getenv(key)

	if len(value) == 0 {
		return def
	} else {
		return strings.TrimSpace(value)
	}
}

const ()

type MonitorConfig struct {
	Path         string
	UpdatePeriod int
}

type AppConfig struct {
	Port          string
	UpdatePeriod  int
	CPUInfoConfig MonitorConfig
	MemInfoConfig MonitorConfig
	UptimeConfig  MonitorConfig
}

func New() (*AppConfig, error) {
	updatePeriod, err := strconv.Atoi(getEnvOrDefault("SSM_PERIOD", "2"))
	if err != nil {
		return nil, err
	}

	// CPU

	updatePeriodCPU, err := strconv.Atoi(getEnvOrDefault("SSM_CPUINFO_PERIOD", strconv.Itoa(updatePeriod)))
	if err != nil {
		return nil, err
	}

	cpuInfoConfig := MonitorConfig{
		Path:         getEnvOrDefault("SSM_CPUINFO_PATH", "/proc/stat"),
		UpdatePeriod: updatePeriodCPU,
	}

	// Mem

	updatePeriodMem, err := strconv.Atoi(getEnvOrDefault("SSM_MEMINFO_PERIOD", strconv.Itoa(updatePeriod)))
	if err != nil {
		return nil, err
	}

	memInfoConfig := MonitorConfig{
		Path:         getEnvOrDefault("SSM_MEMINFO_PATH", "/proc/meminfo"),
		UpdatePeriod: updatePeriodMem,
	}

	// Uptime

	updatePeriodUptime, err := strconv.Atoi(getEnvOrDefault("SSM_UPTIME_PERIOD", strconv.Itoa(updatePeriod)))
	if err != nil {
		return nil, err
	}

	uptimeConfig := MonitorConfig{
		Path:         getEnvOrDefault("SSM_UPTIME_PATH", "/proc/uptime"),
		UpdatePeriod: updatePeriodUptime,
	}

	return &AppConfig{
		Port:          getEnvOrDefault("SSM_PORT", "4442"),
		UpdatePeriod:  updatePeriod,
		CPUInfoConfig: cpuInfoConfig,
		MemInfoConfig: memInfoConfig,
		UptimeConfig:  uptimeConfig,
	}, nil
}

func (ac *AppConfig) String() string {
	out := "SSM config:"
	out += fmt.Sprintf("\nPort: %v", ac.Port)
	out += fmt.Sprintf("\nGlobal update period: %v", ac.UpdatePeriod)
	out += fmt.Sprintf("\nCPU info file path: %v", ac.CPUInfoConfig.Path)
	out += fmt.Sprintf("\nCPU info update period: %v", ac.CPUInfoConfig.UpdatePeriod)
	out += fmt.Sprintf("\nMem info file path: %v", ac.MemInfoConfig.Path)
	out += fmt.Sprintf("\nMem info update period: %v", ac.MemInfoConfig.UpdatePeriod)
	out += fmt.Sprintf("\nUptime file path: %v", ac.UptimeConfig.Path)
	out += fmt.Sprintf("\nUptime update period: %v", ac.UptimeConfig.UpdatePeriod)

	return out
}

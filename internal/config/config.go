package config

import (
	"fmt"
)

const (
	DefaultPeriod     = 2
	DefaultBasePath   = "/proc"
	DefaultCPUPath    = "/stat"
	DefaultMemPath    = "/meminfo"
	DefaultUptimePath = "/uptime"
	DefaultPort       = "4442"
)

type EnvProvider interface {
	GetStringOrDefault(name, def string) string
	GetIntOrDefault(name string, def int) (int, error)
}

type MonitorConfig struct {
	Path         string
	UpdatePeriod int
}

type AppConfig struct {
	Port          string
	UpdatePeriod  int
	Path          string
	CPUInfoConfig MonitorConfig
	MemInfoConfig MonitorConfig
	UptimeConfig  MonitorConfig
}

func New(envProvider EnvProvider) (*AppConfig, error) {
	updatePeriod, err := envProvider.GetIntOrDefault("SSM_PERIOD", DefaultPeriod)
	if err != nil {
		return nil, err
	}

	path := envProvider.GetStringOrDefault("SSM_PATH", DefaultBasePath)

	// CPU

	updatePeriodCPU, err := envProvider.GetIntOrDefault("SSM_CPUINFO_PERIOD", updatePeriod)
	if err != nil {
		return nil, err
	}

	cpuInfoConfig := MonitorConfig{
		Path:         envProvider.GetStringOrDefault("SSM_CPUINFO_PATH", path+DefaultCPUPath),
		UpdatePeriod: updatePeriodCPU,
	}

	// Mem

	updatePeriodMem, err := envProvider.GetIntOrDefault("SSM_MEMINFO_PERIOD", updatePeriod)
	if err != nil {
		return nil, err
	}

	memInfoConfig := MonitorConfig{
		Path:         envProvider.GetStringOrDefault("SSM_MEMINFO_PATH", path+DefaultMemPath),
		UpdatePeriod: updatePeriodMem,
	}

	// Uptime

	updatePeriodUptime, err := envProvider.GetIntOrDefault("SSM_UPTIME_PERIOD", updatePeriod)
	if err != nil {
		return nil, err
	}

	uptimeConfig := MonitorConfig{
		Path:         envProvider.GetStringOrDefault("SSM_UPTIME_PATH", path+DefaultUptimePath),
		UpdatePeriod: updatePeriodUptime,
	}

	return &AppConfig{
		Port:          envProvider.GetStringOrDefault("SSM_PORT", DefaultPort),
		UpdatePeriod:  updatePeriod,
		Path:          path,
		CPUInfoConfig: cpuInfoConfig,
		MemInfoConfig: memInfoConfig,
		UptimeConfig:  uptimeConfig,
	}, nil
}

func (ac *AppConfig) String() string {
	out := "SSM config:"
	out += fmt.Sprintf("\nPort: %v", ac.Port)
	out += fmt.Sprintf("\nGlobal update period: %v", ac.UpdatePeriod)
	out += fmt.Sprintf("\nBase path: %v", ac.Path)
	out += fmt.Sprintf("\nCPU info file path: %v", ac.CPUInfoConfig.Path)
	out += fmt.Sprintf("\nCPU info update period: %v", ac.CPUInfoConfig.UpdatePeriod)
	out += fmt.Sprintf("\nMem info file path: %v", ac.MemInfoConfig.Path)
	out += fmt.Sprintf("\nMem info update period: %v", ac.MemInfoConfig.UpdatePeriod)
	out += fmt.Sprintf("\nUptime file path: %v", ac.UptimeConfig.Path)
	out += fmt.Sprintf("\nUptime update period: %v", ac.UptimeConfig.UpdatePeriod)

	return out
}

package config_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/mishankov/simple-system-monitor/internal/config"
	"github.com/mishankov/testman/assert"
)

func TestDefaultConfig(t *testing.T) {
	env := FakeEnvProvider{defaultOnly: true}

	conf, err := config.New(env)
	assert.NoError(t, err)

	t.Run("test default values", func(t *testing.T) {
		assert.Equal(t, conf.Port, config.DefaultPort)
		assert.Equal(t, conf.UpdatePeriod, config.DefaultPeriod)
		assert.Equal(t, conf.Path, config.DefaultBasePath)
		assert.Equal(t, conf.CPUInfoConfig.Path, config.DefaultBasePath+config.DefaultCPUPath)
		assert.Equal(t, conf.CPUInfoConfig.UpdatePeriod, config.DefaultPeriod)
		assert.Equal(t, conf.MemInfoConfig.Path, config.DefaultBasePath+config.DefaultMemPath)
		assert.Equal(t, conf.MemInfoConfig.UpdatePeriod, config.DefaultPeriod)
		assert.Equal(t, conf.UptimeConfig.Path, config.DefaultBasePath+config.DefaultUptimePath)
		assert.Equal(t, conf.UptimeConfig.UpdatePeriod, config.DefaultPeriod)
	})

	t.Run("test string repr", func(t *testing.T) {
		confStr := conf.String()

		assert.Equal(t, confStr, `SSM config:
Port: 4442
Global update period: 2
Base path: /proc
CPU info file path: /proc/stat
CPU info update period: 2
Mem info file path: /proc/meminfo
Mem info update period: 2
Uptime file path: /proc/uptime
Uptime update period: 2`)
	})
}

func TestErrors(t *testing.T) {
	testCases := []struct {
		key string
		err error
	}{
		{"SSM_PERIOD", errors.New("error")},
		{"SSM_CPUINFO_PERIOD", errors.New("cpu error")},
		{"SSM_MEMINFO_PERIOD", errors.New("mem error")},
		{"SSM_UPTIME_PERIOD", errors.New("uptime error")},
	}

	for _, test := range testCases {
		t.Run(fmt.Sprintf("test %v period error", test.key), func(t *testing.T) {
			conf, err := config.New(FakeEnvProvider{errMap: map[string]error{test.key: test.err}})

			assert.NotNil(t, conf)
			assert.Equal(t, err.Error(), test.err.Error())
		})
	}

}

type FakeEnvProvider struct {
	defaultOnly bool
	err         bool
	intMap      map[string]int
	errMap      map[string]error
}

func (fep FakeEnvProvider) GetStringOrDefault(_, def string) string {
	if fep.defaultOnly {
		return def
	}

	return ""
}

func (fep FakeEnvProvider) GetIntOrDefault(name string, def int) (int, error) {
	val, ok := fep.intMap[name]
	if ok {
		return val, nil
	}

	err, ok := fep.errMap[name]
	if ok {
		return 0, err
	}

	switch {
	case fep.defaultOnly:
		return def, nil
	case fep.err:
		return 0, errors.New("some error")
	}

	return 0, nil
}

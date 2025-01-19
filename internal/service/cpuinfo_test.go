package service_test

import (
	"context"
	"testing"

	"github.com/mishankov/testman/assert"

	"github.com/mishankov/simple-system-monitor/internal/domain/cpuinfo"
	"github.com/mishankov/simple-system-monitor/internal/service"
)

func TestStreamCPULoad(t *testing.T) {
	t.Parallel()

	us := service.NewCPUInfoService(FakeCPULoadRepo{}, 1)
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan []cpuinfo.CPULoad)

	go us.StreamCPULoad(ctx, ch)

	ci := <-ch
	if assert.Equal(t, len(ci), 1) {
		assert.Equal(t, ci[0].ID, "cpu123")
	}

	ci = <-ch
	if assert.Equal(t, len(ci), 1) {
		assert.Equal(t, ci[0].ID, "cpu123")
	}

	ci = <-ch
	if assert.Equal(t, len(ci), 1) {
		assert.Equal(t, ci[0].ID, "cpu123")
	}

	cancel()

	_, ok := <-ch

	assert.Equal(t, ok, false)
}

type FakeCPULoadRepo struct{}

func (fcr FakeCPULoadRepo) GetCPUInfo() ([]cpuinfo.CPUInfo, error) {
	return []cpuinfo.CPUInfo{{ID: "cpu123"}}, nil
}

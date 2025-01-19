package service_test

import (
	"context"
	"testing"

	"github.com/mishankov/testman/assert"

	"github.com/mishankov/simple-system-monitor/internal/domain/meminfo"
	"github.com/mishankov/simple-system-monitor/internal/service"
)

func TestStreamMemInfo(t *testing.T) {
	t.Parallel()

	us := service.NewMemInfoService(FakeMemInfoRepo{}, 1)
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan *meminfo.MemInfo)

	go us.StreamMemInfo(ctx, ch)

	mi := <-ch
	assert.Equal(t, mi.MemTotal, 5)

	mi = <-ch
	assert.Equal(t, mi.MemTotal, 5)

	cancel()

	_, ok := <-ch

	assert.Equal(t, ok, false)
}

type FakeMemInfoRepo struct{}

func (fmr FakeMemInfoRepo) GetMemInfo() (*meminfo.MemInfo, error) {
	return &meminfo.MemInfo{MemTotal: 5, MemFree: 1, MemAvailable: 2}, nil
}

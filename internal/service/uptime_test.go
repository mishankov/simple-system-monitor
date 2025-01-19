package service_test

import (
	"context"
	"testing"

	"github.com/mishankov/simple-system-monitor/internal/domain/uptime"
	"github.com/mishankov/simple-system-monitor/internal/service"
	"github.com/mishankov/testman/assert"
)

func TestStreamUptime(t *testing.T) {
	us := service.NewUptimeService(FakeUptimeRepo{}, 1)
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan *uptime.Uptime)

	go us.StreamUptime(ctx, ch)

	ut := <-ch
	assert.Equal(t, ut.Uptime, 123)

	ut = <-ch
	assert.Equal(t, ut.Uptime, 123)

	cancel()

	_, ok := <-ch

	assert.Equal(t, ok, false)
}

type FakeUptimeRepo struct{}

func (fur FakeUptimeRepo) GetUptime() (*uptime.Uptime, error) {
	return &uptime.Uptime{Uptime: 123}, nil
}

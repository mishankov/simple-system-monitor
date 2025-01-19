package sysinfo

import (
	"testing"

	"github.com/mishankov/testman/assert"

	"github.com/mishankov/simple-system-monitor/internal/testutils"
)

func TestGetUptime(t *testing.T) {
	input := []byte("2244.99 16684.08")
	repo := NewUptimeRepo(testutils.NewFakeFileReader(input))
	data, err := repo.GetUptime()
	assert.NoError(t, err)

	t.Run("test uptime", func(t *testing.T) {
		assert.Equal(t, data.Uptime, float32(2244.99))
	})
}

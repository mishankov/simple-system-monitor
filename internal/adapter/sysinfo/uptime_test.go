package sysinfo

import (
	"testing"

	"github.com/mishankov/simple-system-monitor/internal/testutils"
)

func TestGetUptime(t *testing.T) {
	input := []byte("2244.99 16684.08")
	repo := NewUptimeRepo(testutils.NewFakeFileReader(input))
	data, err := repo.GetUptime()
	testutils.AssertError(t, err)

	t.Run("test uptime", func(t *testing.T) {
		testutils.Assert(t, data.Uptime, float32(2244.99))
	})
}

package sysinfo

import (
	"ssm/internal/testutils"
	"testing"
)

func TestGetUptime(t *testing.T) {
	input := []byte("2244.99 16684.08")
	repo := NewUptimeRepo(testutils.NewFakeFileReader(input))
	data, err := repo.GetUptime()
	if err != nil {
		t.Fatal("Error is not expected")
	}

	t.Run("test uptime", func(t *testing.T) {
		testutils.Assert(t, data.Uptime, float32(2244.99))
	})
}

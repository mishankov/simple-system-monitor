package sysinfo

import (
	"testing"

	"github.com/mishankov/simple-system-monitor/internal/testutils"
	"github.com/mishankov/testman/assert"
)

func TestGetMemInfo(t *testing.T) {
	input := []byte(`MemTotal:        8034976 kB
MemFree:         5207492 kB
MemAvailable:    6676380 kB
Buffers:          108312 kB
Cached:          1621444 kB
SomeOtherData:    101019 kB`)

	repo := NewMemInfoRepo(testutils.NewFakeFileReader(input))

	data, err := repo.GetMemInfo()
	assert.NoError(t, err)

	t.Run("test mem free", func(t *testing.T) {
		assert.Equal(t, data.MemFree, 5207492)
	})

	t.Run("test mem available", func(t *testing.T) {
		assert.Equal(t, data.MemAvailable, 6676380)
	})

	t.Run("test mem total", func(t *testing.T) {
		assert.Equal(t, data.MemTotal, 8034976)
	})
}

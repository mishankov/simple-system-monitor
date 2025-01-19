package fsutils_test

import (
	"os"
	"testing"

	"github.com/mishankov/simple-system-monitor/internal/fsutils"
	"github.com/mishankov/testman/assert"
)

func TestReadData(t *testing.T) {
	path := t.TempDir() + "/file.txt"
	data := []byte("some data")
	err := os.WriteFile(path, data, 0777)
	assert.NoError(t, err)

	fr := fsutils.NewFileReader(path)
	rdata, err := fr.ReadData()

	assert.NoError(t, err)
	assert.Equal(t, string(rdata), string(data))
}

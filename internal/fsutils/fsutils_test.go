package fsutils_test

import (
	"os"
	"testing"

	"github.com/mishankov/simple-system-monitor/internal/fsutils"
	"github.com/mishankov/simple-system-monitor/internal/testutils"
)

func TestReadData(t *testing.T) {
	path := t.TempDir() + "/file.txt"
	data := []byte("some data")
	os.WriteFile(path, data, 0777)

	fr := fsutils.NewFileReader(path)
	rdata, err := fr.ReadData()

	testutils.AssertError(t, err)
	testutils.Assert(t, string(rdata), string(data))
}

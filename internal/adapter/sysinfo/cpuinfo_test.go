package sysinfo

import (
	"strconv"
	"testing"

	"github.com/mishankov/simple-system-monitor/internal/testutils"
)

func TestCPUInfo(t *testing.T) {
	input := []byte(`cpu  231805 945 143933 41984250 205549 0 37459 0 0 0
cpu0 29635 329 18911 5229509 27977 0 5483 0 0 0
cpu1 31541 21 20240 5221500 24229 0 9259 0 0 0
intr 29801338 51 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 5 0 0 0 728 29 0 0 0 0 0 0 0 0 1 335522 142 142 0 0 0 0 0 0 0 0 0 0 0 0 0 0 3 21 995 917 4139423 1537171 2323500 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
ctxt 44303539
btime 1735696829
processes 80117
procs_running 2
procs_blocked 0
softirq 36716593 241 3152731 81484 22133904 337881 0 66331 7962216 270 2981535`)

	repo := CPUInfoRepo{dataReader: testutils.NewFakeFileReader(input)}

	cpuInfos, err := repo.GetCPUInfo()
	testutils.AssertError(t, err)

	t.Run("test CPUs count", func(t *testing.T) {
		if len(cpuInfos) != 2 {
			t.Fatalf("Got %d CPUs want 2", len(cpuInfos))
		}
	})

	t.Run("test CPU ids", func(t *testing.T) {
		for i, cpuInfo := range cpuInfos {
			want := "cpu" + strconv.Itoa(i)
			if cpuInfo.Id != want {
				t.Fatalf("Got CPU id %v want %v", cpuInfo.Id, want)
			}
		}
	})
}

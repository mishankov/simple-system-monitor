package cpuinfo

type CPUInfo struct {
	Id        string
	User      int
	Nice      int
	System    int
	Idle      int
	Iowait    int
	Irq       int
	Softirq   int
	Steal     int
	Guest     int
	GuestNice int
}

type CPUInfoRepo interface {
	GetCPUInfo() ([]CPUInfo, error)
}

type CPULoad struct {
	Id   string  `json:"id"`
	Load float32 `json:"load"`
}

type CPUInfoService interface {
	StreamCPULoad(ch chan []CPULoad)
}

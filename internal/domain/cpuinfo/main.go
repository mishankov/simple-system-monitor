package cpuinfo

type CPUInfo struct {
	Id        string `json:"id"`
	User      int    `json:"user"`
	Nice      int    `json:"nice"`
	System    int    `json:"system"`
	Idle      int    `json:"idle"`
	Iowait    int    `json:"iowait"`
	Irq       int    `json:"irq"`
	Softirq   int    `json:"softirq"`
	Steal     int    `json:"steal"`
	Guest     int    `json:"guest"`
	GuestNice int    `json:"guest_nice"`
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

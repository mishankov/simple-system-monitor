package meminfo

type MemInfo struct {
	MemTotal     int `json:"mem_total"`
	MemFree      int `json:"mem_free"`
	MemAvailable int `json:"mem_available"`
}

type MemInfoRepo interface {
	GetMemInfo() (*MemInfo, error)
}

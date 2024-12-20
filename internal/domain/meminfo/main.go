package meminfo

type MemInfo struct {
	MemTotal     int `json:"mem_total"`
	MemFree      int `json:"mem_free"`
	MemAvailable int `json:"mem_available"`
}

type MemInfoService interface {
	GetMemInfo() (*MemInfo, error)
}

package meminfo

import "context"

type MemInfo struct {
	MemTotal     int `json:"mem_total"`
	MemFree      int `json:"mem_free"`
	MemAvailable int `json:"mem_available"`
}

type Repo interface {
	GetMemInfo() (*MemInfo, error)
}

type Service interface {
	StreamMemInfo(ctx context.Context, ch chan *MemInfo)
}

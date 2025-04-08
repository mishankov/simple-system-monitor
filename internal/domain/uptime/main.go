package uptime

import "context"

type Uptime struct {
	Uptime float32 `json:"uptime"`
}

type Repo interface {
	GetUptime() (*Uptime, error)
}

type Service interface {
	StreamUptime(ctx context.Context, ch chan *Uptime)
}

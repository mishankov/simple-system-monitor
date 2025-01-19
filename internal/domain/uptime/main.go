package uptime

import "context"

type Uptime struct {
	Uptime float32 `json:"uptime"`
}

type UptimeRepo interface {
	GetUptime() (*Uptime, error)
}

type UptimeService interface {
	StreamUptime(ctx context.Context, ch chan *Uptime)
}

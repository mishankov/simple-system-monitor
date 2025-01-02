package uptime

type Uptime struct {
	Uptime float32 `json:"uptime"`
}

type UptimeRepo interface {
	GetUptime() (*Uptime, error)
}

type UptimeService interface {
	StreamUptime(ch chan *Uptime)
}

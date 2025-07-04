package websocket

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mishankov/simple-system-monitor/internal/domain/uptime"
)

type UptimeHandler struct {
	svc uptime.Service
}

func NewUptimeHandler(svc uptime.Service) *UptimeHandler {
	return &UptimeHandler{svc: svc}
}

func (uh *UptimeHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	logger.Infof("%v requests uptime info", req.RemoteAddr)
	defer logger.Info("Stop sending uptime to", req.RemoteAddr)

	ctx, cancel := context.WithCancel(req.Context())
	defer cancel()

	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		logger.Error("Error upgrading to ws:", err)
		return
	}
	defer conn.Close()

	ch := make(chan *uptime.Uptime)
	go uh.svc.StreamUptime(ctx, ch)

	for u := range ch {
		uBytes, _ := json.Marshal(u)
		err := conn.WriteMessage(1, uBytes)
		if err != nil {
			logger.Errorf("Error sending mem info to %v: %v", req.RemoteAddr, err)
			cancel()
		}
	}
}

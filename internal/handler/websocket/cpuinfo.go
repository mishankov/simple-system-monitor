package websocket

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mishankov/simple-system-monitor/internal/domain/cpuinfo"
)

type CPUInfoHandler struct {
	svc cpuinfo.Service
}

func NewCPUInfoHandler(svc cpuinfo.Service) *CPUInfoHandler {
	return &CPUInfoHandler{svc: svc}
}

func (cih *CPUInfoHandler) GetJSONWS(w http.ResponseWriter, req *http.Request) {
	logger.Infof("%v requests CPU info", req.RemoteAddr)
	defer logger.Info("Stop sending cpu info to", req.RemoteAddr)

	ctx, cancel := context.WithCancel(req.Context())
	defer cancel()

	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		logger.Error("Error upgrading to ws:", err)
		return
	}
	defer conn.Close()

	ch := make(chan []cpuinfo.CPULoad)

	go cih.svc.StreamCPULoad(ctx, ch)

	for ci := range ch {
		ciBytes, _ := json.Marshal(ci)
		err := conn.WriteMessage(1, ciBytes)
		if err != nil {
			logger.Errorf("Error sending cpu info to %v: %v", req.RemoteAddr, err)
			cancel()
		}
	}
}

package websocket

import (
	"encoding/json"
	"net/http"

	"github.com/mishankov/simple-system-monitor/internal/domain/cpuinfo"
)

type CPUInfoHandler struct {
	svc cpuinfo.CPUInfoService
}

func NewCPUInfoHandler(svc cpuinfo.CPUInfoService) *CPUInfoHandler {
	return &CPUInfoHandler{svc: svc}
}

func (cih *CPUInfoHandler) GetJsonWS(w http.ResponseWriter, req *http.Request) {
	logger.Infof("%v requests CPU info", req.RemoteAddr)
	defer logger.Info("Stop sending cpu info to", req.RemoteAddr)

	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		logger.Error("Error upgrading to ws:", err)
		return
	}
	defer conn.Close()

	ch := make(chan []cpuinfo.CPULoad)

	go cih.svc.StreamCPULoad(ch)

	for ci := range ch {
		ciBytes, _ := json.Marshal(ci)
		err := conn.WriteMessage(1, ciBytes)
		if err != nil {
			logger.Errorf("Error sending cpu info to %v: %v", req.RemoteAddr, err)
			break
		}
	}
}

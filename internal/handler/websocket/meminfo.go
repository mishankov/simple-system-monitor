package websocket

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mishankov/simple-system-monitor/internal/domain/meminfo"
)

type MemInfoHandler struct {
	svc meminfo.Service
}

func NewMemInfoHandler(svc meminfo.Service) *MemInfoHandler {
	return &MemInfoHandler{svc: svc}
}

func (mif *MemInfoHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	logger.Infof("%v requests mem info", req.RemoteAddr)
	defer logger.Info("Stop sending mem info to", req.RemoteAddr)

	ctx, cancel := context.WithCancel(req.Context())
	defer cancel()

	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		logger.Error("Error upgrading to ws:", err)
		return
	}
	defer conn.Close()

	ch := make(chan *meminfo.MemInfo)
	go mif.svc.StreamMemInfo(ctx, ch)

	for mi := range ch {
		miBytes, _ := json.Marshal(mi)
		err := conn.WriteMessage(1, miBytes)
		if err != nil {
			logger.Errorf("Error sending mem info to %v: %v", req.RemoteAddr, err)
			cancel()
		}
	}
}

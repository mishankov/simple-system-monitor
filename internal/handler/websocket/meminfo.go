package websocket

import (
	"encoding/json"
	"net/http"

	"github.com/mishankov/simple-system-monitor/internal/domain/meminfo"
)

type MemInfoHandler struct {
	svc meminfo.MemInfoService
}

func NewMemInfoHandler(svc meminfo.MemInfoService) *MemInfoHandler {
	return &MemInfoHandler{svc: svc}
}

func (mif *MemInfoHandler) GetJSONWS(w http.ResponseWriter, req *http.Request) {
	logger.Infof("%v requests mem info", req.RemoteAddr)
	defer logger.Info("Stop sending mem info to", req.RemoteAddr)

	ctx := req.Context()

	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		logger.Error("Error upgrading to ws:", err)
		return
	}
	defer conn.Close()

	ch := make(chan *meminfo.MemInfo)
	go mif.svc.StreamMemInfo(ch)

	for {
		doBreak := false
		select {
		case mi := <-ch:
			miBytes, _ := json.Marshal(mi)
			err := conn.WriteMessage(1, miBytes)
			if err != nil {
				logger.Errorf("Error sending mem info to %v: %v", req.RemoteAddr, err)
				doBreak = true
			}
		case <-ctx.Done():
			logger.Infof("Connection of %v closed: %v", req.RemoteAddr, ctx.Err())
			doBreak = true
		}

		if doBreak {
			break
		}
	}
}

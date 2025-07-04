package websocket

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/mishankov/simple-system-monitor/internal/domain/cpuinfo"
	"github.com/mishankov/simple-system-monitor/internal/domain/meminfo"
	"github.com/mishankov/simple-system-monitor/internal/domain/uptime"
)

type CombinedHandler struct {
	cpuSvc    cpuinfo.Service
	memSvc    meminfo.Service
	uptimeSvc uptime.Service
}

func NewCombinedHandler(cpuSvc cpuinfo.Service, memSvc meminfo.Service, uptimeSvc uptime.Service) *CombinedHandler {
	return &CombinedHandler{
		cpuSvc:    cpuSvc,
		memSvc:    memSvc,
		uptimeSvc: uptimeSvc,
	}
}

func (coh *CombinedHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	logger.Infof("%v requests combined info", req.RemoteAddr)
	defer logger.Info("Stop sending combined info to", req.RemoteAddr)

	ctx, cancel := context.WithCancel(req.Context())
	defer cancel()

	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		logger.Error("Error upgrading to ws:", err)
		return
	}
	defer conn.Close()

	messagesCh := make(chan []byte)

	go func() {
		ch := make(chan []cpuinfo.CPULoad)

		go coh.cpuSvc.StreamCPULoad(ctx, ch)

		for ci := range ch {
			ciBytes, _ := json.Marshal(ci)
			messagesCh <- ciBytes
		}
	}()

	go func() {
		ch := make(chan *meminfo.MemInfo)
		go coh.memSvc.StreamMemInfo(ctx, ch)

		for mi := range ch {
			miBytes, _ := json.Marshal(mi)
			messagesCh <- miBytes
		}
	}()

	go func() {
		ch := make(chan *uptime.Uptime)
		go coh.uptimeSvc.StreamUptime(ctx, ch)

		for u := range ch {
			uBytes, _ := json.Marshal(u)
			messagesCh <- uBytes
		}
	}()

	for m := range messagesCh {
		err := conn.WriteMessage(1, m)
		if err != nil {
			logger.Errorf("Error sending data to %v: %v", req.RemoteAddr, err)
			cancel()
		}
	}
}

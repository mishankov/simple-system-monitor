package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"ssm/internal/domain/cpuinfo"
	"ssm/internal/domain/meminfo"
	"ssm/internal/domain/uptime"
)

type CombinedHandler struct {
	cpuSvc    cpuinfo.CPUInfoService
	memSvc    meminfo.MemInfoService
	uptimeSvc uptime.UptimeService
}

func NewCombinedHandler(cpuSvc cpuinfo.CPUInfoService, memSvc meminfo.MemInfoService, uptimeSvc uptime.UptimeService) *CombinedHandler {
	return &CombinedHandler{
		cpuSvc:    cpuSvc,
		memSvc:    memSvc,
		uptimeSvc: uptimeSvc,
	}
}

func (coh *CombinedHandler) GetJsonWS(w http.ResponseWriter, req *http.Request) {
	log.Printf("%v requests combined info", req.RemoteAddr)
	defer log.Println("Stop sending combined info to", req.RemoteAddr)

	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println("Error upgrading to ws:", err)
		return
	}
	defer conn.Close()

	messagesCh := make(chan []byte)

	go func() {
		ch := make(chan []cpuinfo.CPULoad)

		go coh.cpuSvc.StreamCPULoad(ch)

		for ci := range ch {
			ciBytes, _ := json.Marshal(ci)
			messagesCh <- ciBytes
		}
	}()

	go func() {
		ch := make(chan *meminfo.MemInfo)
		go coh.memSvc.StreamMemInfo(ch)

		for mi := range ch {
			miBytes, _ := json.Marshal(mi)
			messagesCh <- miBytes
		}
	}()

	go func() {
		ch := make(chan *uptime.Uptime)
		go coh.uptimeSvc.StreamUptime(ch)

		for u := range ch {
			uBytes, _ := json.Marshal(u)
			messagesCh <- uBytes
		}
	}()

	for m := range messagesCh {
		err := conn.WriteMessage(1, m)
		if err != nil {
			log.Printf("Error sending data to %v: %v", req.RemoteAddr, err)
			break
		}
	}
}

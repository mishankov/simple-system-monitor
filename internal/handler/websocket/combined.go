package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"ssm/internal/domain/cpuinfo"
	"ssm/internal/domain/meminfo"
	"ssm/internal/domain/uptime"
	"sync"
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

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		ch := make(chan []cpuinfo.CPULoad)

		go coh.cpuSvc.StreamCPULoad(ch)

		for ci := range ch {
			ciBytes, _ := json.Marshal(ci)
			err := conn.WriteMessage(1, ciBytes)
			if err != nil {
				log.Printf("Error sending cpu info to %v: %v", req.RemoteAddr, err)
				wg.Done()
				break
			}
		}
	}()

	wg.Add(1)
	go func() {
		ch := make(chan *meminfo.MemInfo)
		go coh.memSvc.StreamMemInfo(ch)

		for mi := range ch {
			miBytes, _ := json.Marshal(mi)
			err := conn.WriteMessage(1, miBytes)
			if err != nil {
				log.Printf("Error sending mem info to %v: %v", req.RemoteAddr, err)
				wg.Done()
				break
			}
		}
	}()

	wg.Add(1)
	go func() {
		ch := make(chan *uptime.Uptime)
		go coh.uptimeSvc.StreamUptime(ch)

		for u := range ch {
			uBytes, _ := json.Marshal(u)
			err := conn.WriteMessage(1, uBytes)
			if err != nil {
				log.Printf("Error sending mem info to %v: %v", req.RemoteAddr, err)
				wg.Done()
				break
			}
		}
	}()

	wg.Wait()
}

package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"ssm/internal/domain/cpuinfo"
	"time"
)

type CPUInfoHandler struct {
	svc cpuinfo.CPUInfoService
}

func NewCPUInfoHandler(svc cpuinfo.CPUInfoService) *CPUInfoHandler {
	return &CPUInfoHandler{svc: svc}
}

func (cih *CPUInfoHandler) GetJsonWS(w http.ResponseWriter, req *http.Request) {
	log.Println("Getting CPU info")

	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println("Error upgrading to ws:", err)
		return
	}

	ch := make(chan []cpuinfo.CPULoad)
	go cih.svc.StreamCPULoad(ch, 2*time.Second)

	for ci := range ch {
		ciBytes, _ := json.Marshal(ci)
		conn.WriteMessage(1, ciBytes)
	}
}

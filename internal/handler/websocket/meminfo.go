package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"ssm/internal/domain/meminfo"
)

type MemInfoHandler struct {
	svc meminfo.MemInfoService
}

func NewMemInfoHandler(svc meminfo.MemInfoService) *MemInfoHandler {
	return &MemInfoHandler{svc: svc}
}

func (mif *MemInfoHandler) GetJsonWS(w http.ResponseWriter, req *http.Request) {
	log.Println("Getting mem info")

	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println("Error upgrading to ws:", err)
		return
	}

	ch := make(chan *meminfo.MemInfo)
	go mif.svc.StreamMemInfo(ch)

	for mi := range ch {
		miBytes, _ := json.Marshal(mi)
		conn.WriteMessage(1, miBytes)
	}
}

package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"ssm/internal/domain/meminfo"
	"time"
)

type MemInfoHandler struct {
	svc meminfo.MemInfoService
}

func NewMemInfoHandler(svc meminfo.MemInfoService) *MemInfoHandler {
	return &MemInfoHandler{svc: svc}
}

func (mif *MemInfoHandler) GetJsonWS(w http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println("Error upgrading to ws:", err)
		return
	}

	mich := make(chan []byte)
	go func() {
		for {
			memInfo, _ := mif.svc.GetMemService()
			memInfoJson, _ := json.Marshal(memInfo)

			mich <- memInfoJson

			time.Sleep(2 * time.Second)
		}
	}()

	for mi := range mich {
		conn.WriteMessage(1, mi)
	}
}

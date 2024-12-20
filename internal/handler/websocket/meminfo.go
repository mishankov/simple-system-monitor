package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"ssm/internal/domain/meminfo"
	"time"
)

type MemInfoHandler struct {
	svc meminfo.MemInfoRepo
}

func NewMemInfoHandler(svc meminfo.MemInfoRepo) *MemInfoHandler {
	return &MemInfoHandler{svc: svc}
}

func (mif *MemInfoHandler) GetJsonWS(w http.ResponseWriter, req *http.Request) {
	log.Println("Getting mem info")

	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println("Error upgrading to ws:", err)
		return
	}

	mich := make(chan []byte)
	go func() {
		for {
			memInfo, _ := mif.svc.GetMemInfo()
			memInfoJson, _ := json.Marshal(memInfo)

			mich <- memInfoJson

			time.Sleep(2 * time.Second)
		}
	}()

	for mi := range mich {
		conn.WriteMessage(1, mi)
	}
}

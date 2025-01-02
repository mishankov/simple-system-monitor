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
	log.Printf("%v requests mem info", req.RemoteAddr)
	defer log.Println("Stop sending mem info to", req.RemoteAddr)

	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println("Error upgrading to ws:", err)
		return
	}
	defer conn.Close()

	ch := make(chan *meminfo.MemInfo)
	go mif.svc.StreamMemInfo(ch)

	for mi := range ch {
		miBytes, _ := json.Marshal(mi)
		err := conn.WriteMessage(1, miBytes)
		if err != nil {
			log.Printf("Error sending mem info to %v: %v", req.RemoteAddr, err)
			break
		}
	}
}

package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"ssm/internal/domain/uptime"
)

type UptimeHandler struct {
	svc uptime.UptimeService
}

func NewUptimeHandler(svc uptime.UptimeService) *UptimeHandler {
	return &UptimeHandler{svc: svc}
}

func (uh *UptimeHandler) GetJsonWS(w http.ResponseWriter, req *http.Request) {
	log.Printf("%v requests uptime info", req.RemoteAddr)

	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println("Error upgrading to ws:", err)
		return
	}

	ch := make(chan *uptime.Uptime)
	go uh.svc.StreamUptime(ch)

	for u := range ch {
		uBytes, _ := json.Marshal(u)
		conn.WriteMessage(1, uBytes)
	}
}

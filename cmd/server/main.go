package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
)

const ASSETS = "public"

func TailMemInfo(mich chan []byte, d time.Duration) {
	for {
		memInfo, _ := GetMemInfo()
		memInfoJson, _ := json.Marshal(memInfo)

		mich <- memInfoJson

		time.Sleep(d)
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func handleMemInfo(w http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Println("Error upgrading to ws:", err)
		return
	}

	mich := make(chan []byte)
	go TailMemInfo(mich, 2*time.Second)

	for mi := range mich {
		conn.WriteMessage(1, mi)
	}
}

func router() chi.Router {
	r := chi.NewRouter()

	r.Get("/meminfo", handleMemInfo)

	r.Handle("/*", http.FileServer(http.Dir(ASSETS)))

	return r
}

func main() {
	r := router()

	log.Printf("Starting server: http://localhost:%v\n", 4442)

	if err := http.ListenAndServe(":"+strconv.Itoa(4442), r); err != nil {
		log.Fatal("Can't start server:", err)
	}
}

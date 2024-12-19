package main

import (
	"log"
	"ssm/internal/adapter/sysinfo"
	"ssm/internal/handler/websocket"
)

func main() {
	memInfoService := sysinfo.NewMemInfoService()
	memInfoHandler := websocket.NewMemInfoHandler(memInfoService)
	server := websocket.NewServer(memInfoHandler, "public", "4442")

	if err := server.Serve(); err != nil {
		log.Println("Error starting server:", err)
	}
}

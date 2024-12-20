package main

import (
	"log"
	"ssm/internal/adapter/sysinfo"
	"ssm/internal/handler/websocket"
)

func main() {
	memInfoService := sysinfo.NewMemInfoRepo()
	memInfoHandler := websocket.NewMemInfoHandler(memInfoService)

	cpuInfoService := sysinfo.NewCPUInfoRepo()
	cpuInfoHandler := websocket.NewCPUInfoHandler(cpuInfoService)

	server := websocket.NewServer(memInfoHandler, cpuInfoHandler, "webapp/build", "4442")

	if err := server.Serve(); err != nil {
		log.Println("Error starting server:", err)
	}
}

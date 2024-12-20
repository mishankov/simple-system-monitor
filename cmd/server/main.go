package main

import (
	"log"
	"ssm/internal/adapter/sysinfo"
	"ssm/internal/handler/websocket"
	"ssm/internal/service"
)

func main() {
	memInfoRepo := sysinfo.NewMemInfoRepo()
	memInfoService := service.NewMemInfoService(memInfoRepo)
	memInfoHandler := websocket.NewMemInfoHandler(memInfoService)

	cpuInfoRepo := sysinfo.NewCPUInfoRepo()
	cpuInfoService := service.NewCPUInfoService(cpuInfoRepo)
	cpuInfoHandler := websocket.NewCPUInfoHandler(cpuInfoService)

	server := websocket.NewServer(memInfoHandler, cpuInfoHandler, "webapp/build", "4442")

	if err := server.Serve(); err != nil {
		log.Println("Error starting server:", err)
	}
}

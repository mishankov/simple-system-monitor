package main

import (
	"embed"
	"log"
	"ssm/internal/adapter/config"
	"ssm/internal/adapter/sysinfo"
	"ssm/internal/handler/websocket"
	"ssm/internal/service"
)

//go:embed all:build
var assets embed.FS

func main() {

	appConfig, err := config.New()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	log.Println(appConfig)

	memInfoRepo := sysinfo.NewMemInfoRepo(appConfig.MemInfoConfig.Path)
	memInfoService := service.NewMemInfoService(memInfoRepo, appConfig.MemInfoConfig.UpdatePeriod)
	memInfoHandler := websocket.NewMemInfoHandler(memInfoService)

	cpuInfoRepo := sysinfo.NewCPUInfoRepo(appConfig.CPUInfoConfig.Path)
	cpuInfoService := service.NewCPUInfoService(cpuInfoRepo, appConfig.CPUInfoConfig.UpdatePeriod)
	cpuInfoHandler := websocket.NewCPUInfoHandler(cpuInfoService)

	server := websocket.NewServer(memInfoHandler, cpuInfoHandler, assets, appConfig.Port)

	if err := server.Serve(); err != nil {
		log.Println("Error starting server:", err)
	}
}

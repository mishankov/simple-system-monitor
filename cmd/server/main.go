package main

import (
	"embed"
	"log"

	"github.com/mishankov/logman/loggers"

	"github.com/mishankov/simple-system-monitor/internal/adapter/sysinfo"
	"github.com/mishankov/simple-system-monitor/internal/config"
	"github.com/mishankov/simple-system-monitor/internal/env"
	"github.com/mishankov/simple-system-monitor/internal/fsutils"
	"github.com/mishankov/simple-system-monitor/internal/handler/websocket"
	"github.com/mishankov/simple-system-monitor/internal/server"
	"github.com/mishankov/simple-system-monitor/internal/service"
)

//go:embed all:build
var assets embed.FS

var logger = loggers.NewDefaultLogger()

func main() {
	env := env.New()

	appConfig, err := config.New(env)
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	logger.Info(appConfig)

	memInfoFileReader := fsutils.NewFileReader(appConfig.MemInfoConfig.Path)
	memInfoRepo := sysinfo.NewMemInfoRepo(memInfoFileReader)
	memInfoService := service.NewMemInfoService(memInfoRepo, appConfig.MemInfoConfig.UpdatePeriod)
	memInfoHandler := websocket.NewMemInfoHandler(memInfoService)

	cpuInfoFileReader := fsutils.NewFileReader(appConfig.CPUInfoConfig.Path)
	cpuInfoRepo := sysinfo.NewCPUInfoRepo(cpuInfoFileReader)
	cpuInfoService := service.NewCPUInfoService(cpuInfoRepo, appConfig.CPUInfoConfig.UpdatePeriod)
	cpuInfoHandler := websocket.NewCPUInfoHandler(cpuInfoService)

	uptimeFileReader := fsutils.NewFileReader(appConfig.UptimeConfig.Path)
	uptimeRepo := sysinfo.NewUptimeRepo(uptimeFileReader)
	uptimeService := service.NewUptimeService(uptimeRepo, appConfig.UptimeConfig.UpdatePeriod)
	uptimeHandler := websocket.NewUptimeHandler(uptimeService)

	combinedHandler := websocket.NewCombinedHandler(cpuInfoService, memInfoService, uptimeService)

	server := server.NewServer(memInfoHandler, cpuInfoHandler, uptimeHandler, combinedHandler, assets, appConfig.Port)

	if err := server.Serve(); err != nil {
		log.Println("Error starting server:", err)
	}
}

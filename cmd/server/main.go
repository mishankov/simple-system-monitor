package main

import (
	"log"
	"ssm/internal/adapter/sysinfo"
	"ssm/internal/handler/http"
)

func main() {
	memInfoService := sysinfo.NewMemInfoService()
	memInfoHandler := http.NewMemInfoHandler(memInfoService)
	server := http.NewServer(memInfoHandler, "public", "4442")

	if err := server.Serve(); err != nil {
		log.Println("Error starting server:", err)
	}
}

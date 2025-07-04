package server

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/mishankov/logman/loggers"

	"github.com/mishankov/simple-system-monitor/internal/handler/websocket"
)

var logger = loggers.NewDefaultLogger()

type Server struct {
	mux  *http.ServeMux
	port string
}

func NewServer(memInfoHandler *websocket.MemInfoHandler, cpuInfoHandler *websocket.CPUInfoHandler, uptimeHandler *websocket.UptimeHandler, combinedHandler *websocket.CombinedHandler, assets embed.FS, userAssetsPath string, port string) *Server {
	mux := http.NewServeMux()

	mux.Handle("/meminfo", memInfoHandler)
	mux.Handle("/cpuinfo", cpuInfoHandler)
	mux.Handle("/uptime", uptimeHandler)
	mux.Handle("/combined", combinedHandler)

	mux.Handle("/user-assets/", http.StripPrefix("/user-assets/", http.FileServer(http.Dir(userAssetsPath))))

	serverRoot, err := fs.Sub(assets, "build")
	if err != nil {
		log.Fatal(err)
	}

	mux.Handle("/", http.FileServer(http.FS(serverRoot)))

	return &Server{mux: mux, port: port}
}

func (s *Server) Serve() error {
	logger.Info("Staring server at", "http://localhost:"+s.port)

	if err := http.ListenAndServe(":"+s.port, s.mux); err != nil {
		return err
	}

	return nil
}

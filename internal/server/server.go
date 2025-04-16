package server

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mishankov/logman/loggers"

	"github.com/mishankov/simple-system-monitor/internal/handler/websocket"
)

var logger = loggers.NewDefaultLogger()

type Server struct {
	router chi.Router
	port   string
}

func NewServer(memInfoHandler *websocket.MemInfoHandler, cpuInfoHandler *websocket.CPUInfoHandler, uptimeHandler *websocket.UptimeHandler, combinedHandler *websocket.CombinedHandler, assets embed.FS, port string) *Server {
	r := chi.NewRouter()

	r.Get("/meminfo", memInfoHandler.GetJSONWS)
	r.Get("/cpuinfo", cpuInfoHandler.GetJSONWS)
	r.Get("/uptime", uptimeHandler.GetJSONWS)
	r.Get("/combined", combinedHandler.GetJSONWS)

	r.Handle("/user-assets/*", http.StripPrefix("/user-assets/", http.FileServer(http.Dir("./user-assets"))))

	serverRoot, err := fs.Sub(assets, "build")
	if err != nil {
		log.Fatal(err)
	}

	r.Handle("/*", http.FileServer(http.FS(serverRoot)))

	return &Server{router: r, port: port}
}

func (s *Server) Serve() error {
	logger.Info("Staring server at", "http://localhost:"+s.port)

	if err := http.ListenAndServe(":"+s.port, s.router); err != nil {
		return err
	}

	return nil
}

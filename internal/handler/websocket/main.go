package websocket

import (
	"embed"
	"io/fs"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	"github.com/mishankov/logman/loggers"
)

var logger = loggers.NewDefaultLogger()

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(_ *http.Request) bool { return true },
}

type Server struct {
	router chi.Router
	port   string
}

func NewServer(memInfoHandler *MemInfoHandler, cpuInfoHandler *CPUInfoHandler, uptimeHandler *UptimeHandler, combinedHandler *CombinedHandler, assets embed.FS, port string) *Server {
	r := chi.NewRouter()

	r.Get("/meminfo", memInfoHandler.GetJsonWS)
	r.Get("/cpuinfo", cpuInfoHandler.GetJsonWS)
	r.Get("/uptime", uptimeHandler.GetJsonWS)
	r.Get("/combined", combinedHandler.GetJsonWS)

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

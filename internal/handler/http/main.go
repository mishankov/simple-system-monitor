package http

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Server struct {
	router chi.Router
	port   string
}

func NewServer(memInfoHandler *MemInfoHandler, assets string, port string) *Server {
	r := chi.NewRouter()

	r.Get("/meminfo", memInfoHandler.GetJsonWS)
	r.Handle("/*", http.FileServer(http.Dir(assets)))

	return &Server{router: r, port: port}
}

func (s *Server) Serve() error {
	log.Println("Staring server at", "http://localhost:"+s.port)

	if err := http.ListenAndServe(":"+s.port, s.router); err != nil {
		return err
	}

	return nil
}

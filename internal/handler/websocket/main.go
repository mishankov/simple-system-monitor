package websocket

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/mishankov/logman/loggers"
)

var logger = loggers.NewDefaultLogger()

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(_ *http.Request) bool { return true },
}

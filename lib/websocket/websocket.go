package websocket

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type WSHandler struct {
	conn       *websocket.Conn
	writeMutex sync.Mutex
}

func CreateWebsocketHandler(w http.ResponseWriter, r *http.Request) (*WSHandler, error) {
	// TODO use pool
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}
	return &WSHandler{
		conn:       conn,
		writeMutex: sync.Mutex{},
	}, nil
}

func (h *WSHandler) Read() (string, error) {
	_, message, err := h.conn.ReadMessage()
	return string(message), err
}

func (h *WSHandler) Write(message string) error {
	h.writeMutex.Lock()
	defer h.writeMutex.Unlock()
	err := h.conn.WriteMessage(websocket.TextMessage, []byte(message))
	return err
}

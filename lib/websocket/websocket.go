package websocket

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type WSHandler struct {
	conn *websocket.Conn
}

func CreateWebsocketHandler(w http.ResponseWriter, r *http.Request) (*WSHandler, error) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}
	return &WSHandler{conn}, nil
}

func (h *WSHandler) Read() (string, error) {
	_, message, err := h.conn.ReadMessage()
	return string(message), err
}

func (h *WSHandler) Write(message string) error {
	return h.conn.WriteMessage(websocket.TextMessage, []byte(message))
}

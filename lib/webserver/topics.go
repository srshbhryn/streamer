package webserver

// import (
// 	"net/http"
// 	"sync"

// 	"github.com/gorilla/websocket"
// )

// type TopicHandler struct {
// 	receiver <-chan string
// 	topic    []string
// }

// func CreateTopicHandler(w http.ResponseWriter, r *http.Request) (*Handler, error) {
// 	conn, err := wsupgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Handler{conn, &sync.Mutex{}}, nil
// }

// func (h *Handler) Read() (string, error) {
// 	_, message, err := h.conn.ReadMessage()
// 	return string(message), err
// }

// func (h *Handler) Write(message string) error {
// 	return h.conn.WriteMessage(websocket.TextMessage, []byte(message))
// }

package streamer

import (
	"strings"
	"sync"
)

type ReaderWriter interface {
	Read() (string, error)
	Write(string) error
}

type Reader interface {
	Read() (string, error)
}

type client_t ReaderWriter
type source_t Reader

var source source_t
var clients map[client_t]map[string]struct{}
var clientsMutex sync.Mutex

func Init(r Reader) {
	source = r
	clients = make(map[client_t]map[string]struct{}, 2^16)
	clientsMutex = sync.Mutex{}
}

func Add(rw ReaderWriter) {
	if _, ok := clients[client_t(rw)]; ok {
		return
	}
	clientsMutex.Lock()
	defer clientsMutex.Unlock()
	clients[client_t(rw)] = make(map[string]struct{}, 32)
	go reader(rw)
}

func Run() {
	for {
		func() {
			msgAngTopicStr, _ := source.Read()
			msgAngTopic := strings.Split(msgAngTopicStr, ",")
			topic := msgAngTopic[0]
			clientsMutex.Lock()
			defer clientsMutex.Unlock()
			for c, topics := range clients {
				_, ok := topics[topic]
				if !ok {
					continue
				}
				go c.Write(msgAngTopicStr)
			}
		}()
	}
}

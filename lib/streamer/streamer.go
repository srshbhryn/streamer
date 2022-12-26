package streamer

import (
	"strings"
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
var clients map[client_t]struct{}
var topicClients map[string]map[client_t]struct{}

func Init(r Reader) {
	source = r
	clients = make(map[client_t]struct{}, 2^16)
	topicClients = make(map[string]map[client_t]struct{})
	for topic := range config.topics {
		topicClients[topic] = make(map[client_t]struct{}, 2048)
	}
}

func Add(rw ReaderWriter) {
	if _, ok := clients[client_t(rw)]; ok {
		return
	}
	clients[client_t(rw)] = struct{}{}
	go reader(rw)
}

func Run() {
	for {
		msgAngTopicStr, _ := source.Read()
		msgAngTopic := strings.Split(msgAngTopicStr, ",")
		topic := msgAngTopic[0]
		clients, ok := topicClients[topic]
		if !ok {
			continue
		}
		for c := range clients {
			go c.Write(msgAngTopicStr)
		}
	}
}

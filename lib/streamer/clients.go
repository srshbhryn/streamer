package streamer

import (
	"strings"
)

func isTopicValid(topic string) bool {
	_, ok := config.topics[topic]
	return ok
}

func reader(rw ReaderWriter) {
	for {
		commandAndArgs, err := rw.Read()
		if err != nil {
			return
		}
		commandAndArgsSlice := strings.Split(commandAndArgs, ",")
		command := commandAndArgsSlice[0]
		switch command {
		case "subscribe":
			if len(commandAndArgsSlice) != 2 {
				continue
			}
			topic := commandAndArgsSlice[1]
			if !isTopicValid(topic) {
				continue
			}
			if _, ok := topicClients[topic][client_t(rw)]; !ok {
				topicClients[topic][client_t(rw)] = struct{}{}
			}
		case "unsubscribe":
			if len(commandAndArgsSlice) != 2 {
				continue
			}
			topic := commandAndArgsSlice[1]
			if !isTopicValid(topic) {
				continue
			}
			if _, ok := topicClients[topic][client_t(rw)]; ok {
				delete(topicClients[topic], client_t(rw))
			}
		}
	}
}

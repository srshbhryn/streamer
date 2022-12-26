package streamer

import (
	"fmt"
	"strings"
)

func isTopicValid(topic string) bool {
	_, ok := config.topics[topic]
	return ok
}

func reader(rw ReaderWriter) {
	for {
		func() {
			commandAndArgs, err := rw.Read()
			fmt.Println(commandAndArgs)
			if err != nil {
				fmt.Println(err)
				return
			}
			commandAndArgsSlice := strings.Split(commandAndArgs, ",")
			command := commandAndArgsSlice[0]
			switch command {
			case "subscribe":
				if len(commandAndArgsSlice) != 2 {
					return
				}
				topic := commandAndArgsSlice[1]
				if !isTopicValid(topic) {
					return
				}

				clientsMutex.Lock()
				defer clientsMutex.Unlock()
				topics, _ := clients[client_t(rw)]
				topics[topic] = struct{}{}

			case "unsubscribe":
				if len(commandAndArgsSlice) != 2 {
					return
				}
				topic := commandAndArgsSlice[1]
				if !isTopicValid(topic) {
					return
				}

				clientsMutex.Lock()
				defer clientsMutex.Unlock()
				topics, _ := clients[client_t(rw)]
				delete(topics, topic)
			}
		}()
	}
}

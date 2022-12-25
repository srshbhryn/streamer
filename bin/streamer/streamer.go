package main

import (
	"github.com/srshbhryn/streamer/lib/webserver"
)

func main() {
	webserver.Create()
	err := webserver.Run()
	if err != nil {
		panic(err)
	}
}

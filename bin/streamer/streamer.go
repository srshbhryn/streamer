package main

import (
	"github.com/srshbhryn/streamer/lib/pullers"
	"github.com/srshbhryn/streamer/lib/streamer"
	"github.com/srshbhryn/streamer/lib/webserver"
)

func main() {
	webserver.Create()
	p := pullers.CreateMock()
	streamer.Init(p)
	go streamer.Run()
	err := webserver.Run()
	if err != nil {
		panic(err)
	}
}

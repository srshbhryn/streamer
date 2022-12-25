package webserver

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/srshbhryn/streamer/lib/websocket"
)

var r *gin.Engine

func Create() {
	if !config.debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r = gin.Default()

	if config.debug {
		r.LoadHTMLFiles("./static/index.html")
		r.GET("/", func(c *gin.Context) {
			c.HTML(200, "index.html", nil)
		})
		r.Static("/static", "./static")
	}

	r.GET("/getTopics", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H(map[string]any{"code": 0, "topic": []string{"wsd", "xsdasd", "aaaa"}}))
	})

	r.GET("/ws", func(c *gin.Context) {
		handler, err := websocket.CreateWebsocketHandler(
			c.Writer,
			c.Request,
		)
		if err != nil {
			c.JSON(400, "")
		}
		wg := sync.WaitGroup{}
		wg.Add(2)
		go func() {
			for {
				msg, err := handler.Read()
				if err != nil {
					wg.Done()
					wg.Done()
					return
				}
				fmt.Println(msg)
			}
		}()
		go func() {
			for {
				msg := "HeartBeat"
				err := handler.Write(msg)
				if err != nil {
					wg.Done()
					wg.Done()
					return
				}
				time.Sleep(time.Second)
			}
		}()
		wg.Wait()
	})

}

func Run() error {
	return r.Run(":" + "8080")
}

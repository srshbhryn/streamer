package webserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/srshbhryn/streamer/lib/streamer"
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
		streamer.Add(handler)
	})

}

func Run() error {
	return r.Run(":" + "8080")
}

package handlers

import (
	"github.com/869413421/wechatbot/gtp"
	"github.com/gin-gonic/gin"
	"strings"
)

func StartOpenApi() {

	r := gin.Default()
	r.GET("/open", func(c *gin.Context) {

		msg := c.Query("msg")
		if len(msg) == 0 {
			c.String(500, "failed... msg cannot be empty")
		}
		completions, err := gtp.Completions(msg)
		completions = strings.TrimSpace(completions)
		completions = strings.Trim(completions, "\n")
		if err != nil {
			c.String(500, "failed... err:%v", err)
		} else {
			c.String(200, "%v", completions)
		}
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

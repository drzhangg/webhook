package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	log.Println("starting...")
	r := gin.Default()

	fmt.Println("web服务成功启动")

	r.GET("/webhook", func(c *gin.Context) {
		fmt.Println("接收到webhook请求")
	})

	r.POST("/post", func(c *gin.Context) {
		fmt.Println("header:", c.Request.Header)
		fmt.Println("Body:", c.Request.Body)
	})

	r.GET("/info", func(c *gin.Context) {
		username := c.Query("username")

		if username != "" {
			c.JSON(200, map[string]interface{}{
				"name":    username,
				"message": "hello," + username,
			})
			return
		}
	})

	r.Run(":8989")
}

package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	fmt.Println("web服务成功启动")

	r.GET("/webhook", func(c *gin.Context) {
		fmt.Println("接收到webhook请求")
	})


	r.POST("/post", func(c *gin.Context) {
		fmt.Println("header:",c.Request.Header)
		fmt.Println("Body:",c.Request.Body)
	})


	r.Run(":8989")
}

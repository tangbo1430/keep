package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个默认的 Gin 引擎
	router := gin.Default()

	// 定义路由，返回 "Hello, World!"
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	router.GET("/export", Export)

	// 启动 Gin 服务，监听 8080 端口
	router.Run(":8080")
}

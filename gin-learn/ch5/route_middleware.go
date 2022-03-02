package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	// 实例化对象
	route := gin.New()
	route.Use(gin.Logger(), gin.Recovery())
	v1 := route.Group("v1").Use(middlewareAuth)
	{
		v1.GET("/ping", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"hello": "isMe",
			})
		})
	}
	{
		println("test")
		{
			println("test inline")
		}
	}
	route.POST("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// 启动服务
	err := route.Run()
	if err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func middlewareAuth(c *gin.Context) {
	start := time.Now()
	c.Set("a", 1)
	c.Next()
	println("runtime:", time.Since(start))
}

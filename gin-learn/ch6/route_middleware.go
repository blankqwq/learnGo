package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
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
	err := route.Run()
	if err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func middlewareAuth(c *gin.Context) {
	v := c.GetHeader("x-token")
	if v == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg": "must login",
		})
		c.Abort()
		return
	}
	c.Next()
}

package ch1

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// 实例化对象
	route := gin.Default()
	route.GET("/query", func(c *gin.Context) {
		name := c.DefaultQuery("name", "hello")
		id := c.Query("id")
		c.JSON(200, gin.H{
			"id":   id,
			"name": name,
		})
	})
	route.GET("/post", func(c *gin.Context) {
		name := c.DefaultPostForm("name", "hello")
		id := c.PostForm("id")
		c.JSON(200, gin.H{
			"id":   id,
			"name": name,
		})
	})
	// 启动服务
	err := route.Run()
	if err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

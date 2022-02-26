package ch1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 实例化对象
	route := gin.Default()
	route.GET("/:name/:id", func(c *gin.Context) {
		name := c.Param("name")
		id := c.Param("id")

		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"id":   id,
		})
	})
	// 启动服务
	err := route.Run()
	if err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

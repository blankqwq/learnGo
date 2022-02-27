package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 实例化对象
	route := gin.Default()
	route.GET("/json", testJson)
	// 启动服务
	err := route.Run()
	if err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

type User struct {
	Name string `json:"user"`
	Age  int
}

func testJson(c *gin.Context) {
	var user User
	user.Name = "test"
	user.Age = 20
	//{
	//	"user": "test",
	//	"Age": 20
	//}
	c.JSON(http.StatusOK, user)
}

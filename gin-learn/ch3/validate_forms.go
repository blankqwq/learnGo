package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 实例化对象
	route := gin.Default()
	route.POST("/json", testJson)
	route.POST("/register", register)
	// 启动服务
	err := route.Run()
	if err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func register(c *gin.Context) {
	var register RegisterForm
	if err := c.ShouldBindJSON(&register); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, register)
}

type User struct {
	Username string `json:"username" binding:"required,min=3,max=10" `
	Password int    `json:"password" binding:"required" `
}

type RegisterForm struct {
	Username      string `json:"username" binding:"required,min=3,max=10" `
	Age           int32  `json:"age" binding:"required" `
	Password      string `json:"password" binding:"required" `
	CheckPassword string `json:"check_password" binding:"required,eqfield=Password" `
}

func testJson(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

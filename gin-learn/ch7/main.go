package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	route := gin.Default()
	route.LoadHTMLFiles("templates/index.tmpl")
	route.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "hello world",
		})
	})
	err := route.Run()
	if err != nil {
		panic(err)
	}
}

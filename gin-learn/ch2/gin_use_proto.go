package main

import (
	"github.com/gin-gonic/gin"
	"gorpc_study/gin-learn/ch2/proto"
	"net/http"
)

func main() {
	// 实例化对象
	route := gin.Default()
	route.GET("/proto", testProto)
	// 启动服务
	err := route.Run()
	if err != nil {
		panic(err)
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func testProto(c *gin.Context) {
	// 返回的原始字符串
	c.ProtoBuf(http.StatusOK, proto.User{Name: "asd", Ohter: []int32{1, 23}})
}

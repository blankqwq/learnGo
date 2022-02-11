package main

import (
	"fmt"
	//"net/http"
)

//func handle(resp http.ResponseWriter,req *http.Request){
//	resp.Write([]byte(fmt.Sprintf("hello world")))
//}
func handle(c *Context)error{
	println("helloworld")
	c.W.Write([]byte(fmt.Sprintf("hello world")))
	return nil
}

func login(c *Context)error{
	c.HttpOk(struct {
		Data string `json:"data"`
	}{
		"ok",
	},"success")
	return nil
}


func test(c *Context)error{
	println(111222)
	c.HttpOk(struct {
		Data string `json:"data"`
	}{
		fmt.Sprintf("%v",c.RouteParameter),
	},"success")
	return nil
}
// 封装server
func main(){
	 //http.HandleFunc("/", handle)
	//err := http.ListenAndServe(":8080", nil)
	s:=NewServer("test",NewBaseHandle())
	s.Get("/*",handle,&TestFilter{})
	s.Get("/test",login)
	s.Get("/user/:test",test)
	go func() {
		s.Run(":8080")
	}()
}
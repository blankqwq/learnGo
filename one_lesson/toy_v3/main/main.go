package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
	//"net/http"
)

//func handle(resp http.ResponseWriter,req *http.Request){
//	resp.Write([]byte(fmt.Sprintf("hello world")))
//}
func handle(c *Context) error {
	println("helloworld")
	c.W.Write([]byte(fmt.Sprintf("hello world")))
	return nil
}

func login(c *Context) error {
	c.HttpOk(struct {
		Data string `json:"data"`
	}{
		"ok",
	}, "success")
	return nil
}

func test(c *Context) error {
	println(111222)
	c.HttpOk(struct {
		Data string `json:"data"`
	}{
		fmt.Sprintf("%v", c.RouteParameter),
	}, "success")
	return nil
}

// 封装server
func main() {
	//http.HandleFunc("/", handle)
	//err := http.ListenAndServe(":8080", nil)
	s := NewServer("test", NewBaseHandle())
	s.Get("/*", handle, &TestFilter{})
	s.Get("/test", login)
	s.Get("/user/:test", test)
	signalsChain := make(chan os.Signal, 1)
	signal.Notify(signalsChain, ShutDownSignals...)
	go func(hooks ...Hook) {
		select {
		case <-signalsChain:
			// 执行hook
			time.AfterFunc(time.Minute*10, func() {
				fmt.Printf("shutdown  timeout\n")
				os.Exit(-1)
			})
			ctx := context.Background()
			for _, h := range hooks {
				c, cancel := context.WithTimeout(ctx, time.Second*20)
				errs := h(c)
				if errs != nil {
					// 执行
					fmt.Printf("hooks return err: %s", errs.Error())
				}
				cancel()
			}
			//正常推出
			os.Exit(0)
		}
	}(RejectRequestHook, BuildShutdownHook(s))
	s.Run(":8080")

}

package proxy

import (
	"gorpc_study/rpc_origin_plus/handler"
	"net/rpc"
)

type HelloService interface {
	Hello(req string, resp *string) error
}

func ServerProxy(hello HelloService,name string) error  {
	println("register:"+handler.HelloServiceName)
	return rpc.RegisterName(name, hello)
}

//自动生成！new ...Servie => call name => return


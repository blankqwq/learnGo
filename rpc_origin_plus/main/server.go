package main

import (
	"gorpc_study/rpc_origin_plus/handler"
	"gorpc_study/rpc_origin_plus/proxy"
	"gorpc_study/rpc_origin_plus/service"
	"net"
	"net/rpc"
)

func main() {
	listen, _ := net.Listen("tcp", "0.0.0.0:8082")
	_ = proxy.ServerProxy(&service.HelloService{}, handler.HelloServiceName)
	conn, _ := listen.Accept()
	rpc.ServeConn(conn)
}

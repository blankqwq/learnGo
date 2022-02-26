package main

import (
	"net"
	"net/rpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(req string, resp *string) error {
	*resp = "hello," + req
	return nil
}

func (s *HelloService) Hello2(req string, resp *string) error {
	*resp = "hello2," + req
	return nil
}

func main() {
	listen, _ := net.Listen("tcp", "0.0.0.0:8082")
	_ = rpc.RegisterName("Hello", &HelloService{})
	conn, _ := listen.Accept()
	rpc.ServeConn(conn)
}

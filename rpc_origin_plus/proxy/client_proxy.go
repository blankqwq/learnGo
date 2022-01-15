package proxy

import (
	"gorpc_study/rpc_origin_plus/handler"
	"net/rpc"
)

type ClientStub struct {
	client *rpc.Client
}

func NewClient(proto,address string) ClientStub  {
	client,err :=rpc.Dial("tcp","localhost:8082")
	if err !=nil {
		panic(err)
	}
	return ClientStub{client: client}
}

func (c *ClientStub)Hello(name string) string  {
	var res string
	if err1 := c.client.Call(handler.HelloServiceName+".Hello",name,&res) ;err1!=nil {
		panic(err1)
	}
	return res
}
package main

import (
	"fmt"
	"google.golang.org/grpc"
	"gorpc/grpc_stream/proto"
	"net"
	"time"
)

const HOST = "0.0.0.0:8082"

type Service struct {
}

func (s *Service) GetStream(req *proto.StreamReqData, p proto.Greeter_GetStreamServer) error {
	fmt.Println("get data:" + req.Data)
	i := 1
	for {
		_ = p.Send(&proto.StreamResData{Data: fmt.Sprintf("current_time:%d", time.Now().Unix())})
		if i > 10 {
			break
		}
		i++
	}
	return nil
}

func (s *Service) PutStream(p proto.Greeter_PutStreamServer) error {
	for {
		r, err := p.Recv()
		if err != nil {
			break
		}
		fmt.Println("get:" + r.Data)
	}
	fmt.Println("end_get data")
	return nil

}

func (s *Service) AllStream(p proto.Greeter_AllStreamServer) error {
	for i := 1; i <= 10; i++ {
		r, err := p.Recv()
		if err != nil {
			fmt.Println(err)
			break;
		}
		fmt.Println("get:"+r.Data)
		_ = p.Send(&proto.StreamResData{Data: "i'm ok" + r.Data})
	}
	return nil
}

func main() {
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &Service{})
	l, err := net.Listen("tcp", HOST)
	if err != nil {
		panic(err)
	}
	err = g.Serve(l)
	if err != nil {
		panic(err)
	}
}

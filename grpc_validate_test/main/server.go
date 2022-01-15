package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"gorpc/grpc_validate_test/proto"
	"net"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, req *proto.Person) (*proto.HelloReply, error) {
	md,_:=metadata.FromIncomingContext(ctx)
	fmt.Println(md)
	return &proto.HelloReply{Message: "hello," + req.Email}, nil
}

type Validate interface {
	ValidateAll() error
}
func Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error){
	fmt.Println("收到一个请求")
	// 校验
	if r,ok:=req.(Validate);ok {
		if err:=r.ValidateAll();err!=nil {
			return nil,status.Error(codes.InvalidArgument,err.Error())
		}
	}
	return handler(ctx,req)
}
func main()  {
	opt:=grpc.UnaryInterceptor(Interceptor)
	g:= grpc.NewServer(opt)
	proto.RegisterGreeterServer(g,&Server{})
	l,err:=net.Listen("tcp","0.0.0.0:8082")
	if err!=nil {
		panic(err)
	}
	err = g.Serve(l)
	if err!=nil {
		panic(err)
	}
}

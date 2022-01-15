package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"gorpc/grpc_test/proto"
	"net"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	//md, _ := metadata.FromIncomingContext(ctx)
	//fmt.Println(md)
	return &proto.HelloReply{Message: "hello," + req.Name}, nil
}
func Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	fmt.Print("收到一个请求,metadata")
	mt, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return resp,status.Error(codes.Unauthenticated,"未认证")
	}
	if val, ok := mt["app-id"];ok && val[0]=="name"{
		fmt.Println("登录成功",val)
	}
	fmt.Println(mt)
	return handler(ctx, req)
}
func main() {
	opt := grpc.UnaryInterceptor(Interceptor)
	g := grpc.NewServer(opt)
	proto.RegisterGreeterServer(g, &Server{})
	l, err := net.Listen("tcp", "0.0.0.0:8082")
	if err != nil {
		panic(err)
	}
	err = g.Serve(l)
	if err != nil {
		panic(err)
	}
}

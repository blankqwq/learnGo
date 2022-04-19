package main

import (
	"context"
	"fmt"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorpc_study/grpc_retry/proto"
	"time"
)

func main() {

	var opts []grpc.DialOption

	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(
		// 全局配置
		[]grpc_retry.CallOption{
			grpc_retry.WithMax(3), grpc_retry.WithPerRetryTimeout(1 * time.Second), grpc_retry.WithCodes(codes.Unknown),
		}...)))
	conn, err := grpc.Dial("localhost:8082", opts...)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	c := proto.NewGreeterClient(conn)
	r, err := c.SayHello(ctx, &proto.Person{Id: 0})
	if err != nil {
		err1, _ := status.FromError(err)
		fmt.Println(err1.Code(), "----", err1.Message())
		panic(err)
	}
	fmt.Println(r.Message)
}

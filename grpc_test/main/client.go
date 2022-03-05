package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"gorpc_study/grpc_test/proto"
)

func main() {
	conn, err := grpc.Dial("localhost:8082", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	c := proto.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "world"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}

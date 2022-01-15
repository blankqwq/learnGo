package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"gorpc/grpc_validate_test/proto"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:8082", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	ctx,cancel := context.WithTimeout(context.Background(), time.Second*2)
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

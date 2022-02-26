package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorpc_study/proto_test/proto"
)

func main() {
	con, err := grpc.Dial("localhost:8083", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer con.Close()
	c := proto.NewGreeterClient(con)
	r, err := c.Ping(context.Background(), &emptypb.Empty{})
	if err != nil {
		panic(err)
	}
	//proto.Pong{Test: &proto.Pong_TestIn{Name: "hahah"}, Info: "asd", Timestamp: timestamppb.New(time.Now())}
	fmt.Println(r.Info, r.Test.Name)
}

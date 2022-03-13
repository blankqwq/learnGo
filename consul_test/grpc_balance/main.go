package main

import (
	"context"
	"fmt"
	"gorpc_study/consul_test/grpc_balance/proto"
	"log"
	_ "time"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(
		"consul://127.0.0.1:9001/user-srv?wait=14s&tag=python",
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	for n := 0; n < 10; n++ {
		print("=========================\n")
		user := proto.NewUserClient(conn)
		rsp, err := user.GetUserList(context.Background(), &proto.PageInfo{
			Page:  1,
			Limit: 2,
		})
		if err != nil {
			panic(err)
		}
		for i, d := range rsp.Data {
			fmt.Println(i, d)
		}
	}

}

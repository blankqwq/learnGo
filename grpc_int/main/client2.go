package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"gorpc/grpc_test/proto"
	"time"
)

func Requset(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	now := time.Now()
	fmt.Println(now)
	md := metadata.New(map[string]string{
		"app-id": "name",
		"ddd":    "ddd",
	})
	ctx = metadata.NewOutgoingContext(ctx, md)
	err := invoker(ctx, method, req, reply, cc, opts...)
	fmt.Println(time.Now().Unix() - now.Unix())
	return err
}

type customer struct{}

func (c customer) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"app-id": "name",
		"ddd":    "ddd",
	},nil
}
func (c customer) RequireTransportSecurity() bool {
	return false
}
func main() {
	var opts []grpc.DialOption
	grpc.WithPerRPCCredentials(customer{})
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithUnaryInterceptor(Requset))
	conn, err := grpc.Dial("localhost:8082", opts...)
	if err != nil {
		panic(err)
	}
	md := metadata.New(map[string]string{
		"asd": "sdad",
		"sad": "asd",
	})
	//md:=metadata.Pairs("1","2","3","4")
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	c := proto.NewGreeterClient(conn)
	r, err := c.SayHello(ctx, &proto.HelloRequest{Name: "world"})
	if err != nil {
		panic(err)
	}
	fmt.Println(r.Message)
}

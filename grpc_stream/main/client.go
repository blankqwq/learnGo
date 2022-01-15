package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"gorpc/grpc_stream/proto"
)

func main()  {
	conn,err:=grpc.Dial("localhost:8082",grpc.WithInsecure())
	if err!=nil {
		panic(err)
	}
	service := proto.NewGreeterClient(conn)
	r,err :=service.GetStream(context.Background(),&proto.StreamReqData{Data: "hello"})
	if err!=nil {
		panic(err)
	}
	for  {
		data,e:=r.Recv()
		if e!=nil {
			break;
		}
		fmt.Println(data.Data)
	}

	p, err := service.PutStream(context.Background())
	for i:=1;i<=10;i++{
		_ = p.Send(&proto.StreamReqData{Data: fmt.Sprintf("hello,i am user %d", i)})
	}
	fmt.Println("send ok")

	p2, err := service.AllStream(context.Background())
	if err!=nil {
		panic(err)
	}
	_ = p2.Send(&proto.StreamReqData{Data: "oh..ok"})
	r2, err := p2.Recv()
	if err!=nil {
		panic(err)
	}
	fmt.Println("rec:"+r2.Data)
}
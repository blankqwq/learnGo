package main

import (
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct {
}

func (s *HelloService) Hello(req string, resp *string) error {
	*resp = "hello," + req
	return nil
}


func main()  {
	_ = rpc.RegisterName("Hello", &HelloService{})
	http.HandleFunc("/jsonrpc", func(writer http.ResponseWriter, request *http.Request) {
		var conn io.ReadWriteCloser = struct{
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: request.Body,
			Writer: writer,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})
	http.ListenAndServe("0.0.0.0:8082",nil)
}
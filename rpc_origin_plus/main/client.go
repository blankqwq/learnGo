package main

import "gorpc_study/rpc_origin_plus/proxy"

func main() {
	client := proxy.NewClient("tcp", "localhost:8082")
	print(client.Hello("world"))
}

package main

import "net/rpc"

func main() {
	client, err := rpc.Dial("tcp", "localhost:8082")
	if err != nil {
		panic(err)
	}
	res := ""
	if err1 := client.Call("Hello.Hello", "world", &res); err1 != nil {
		panic(err1)
	}
	print(res)
}

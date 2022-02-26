## http服务的JSON RPC

> 基于json协议的rpc服务调用

依托于http服务

```golang
_ = rpc.RegisterName("Hello", &HelloService{})


err := rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
		if err != nil {
			panic(err)
			return
		}
```

调用方式居然还是=>rpc.


```python
from jsonrpclib.jsonrpc import ServerProxy

# if __name__=='__main__':
#     server = ServerProxy("http://localhost:8082")
#     print(server.add(1, 2))


if __name__=='__main__':
    server = ServerProxy("http://localhost:8082/jsonrpc")
    # 好家伙
    print(server.Hello.Hello("assd"))
```

```shell script
protoc -I . hello.proto --go_out=plugins=grpc:.

protoc -I . hello.proto --go_out=plugins=grpc:. --validate_out=lang=go:.
```
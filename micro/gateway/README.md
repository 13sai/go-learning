

```
protoc --proto_path=/Users/wangzetao/Downloads/chrome/googleapis-master/ -I ./proto/ --go_out=plugins=grpc:proto --grpc-gateway_out=logtostderr=true:proto/. proto/hello.proto
```
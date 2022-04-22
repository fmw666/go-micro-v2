```go
go get github.com/micro/go-micro/v2

// 需要到下载位置进行 `go build` 获取可执行文件
go get github.com/golang/protobuf
```

```sh
git clone https://github.com/zserge/protoc-gen-micro
cd protoc-gen-micro
go build
# protoc-gen-micro.exe
```

+ protoc 下载

https://github.com/protocolbuffers/protobuf/releases


```s
protoc --proto_path=. --micro_out=. --go_out=. userModel.proto
protoc --proto_path=. --micro_out=. --go_out=. userService.proto
```
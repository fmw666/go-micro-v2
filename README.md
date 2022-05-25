## 微服务应用代码示例

<br>

<div align="right">
    <a href="https://github.com/fmw666/microservice-code-sample/tree/master#-%E5%88%86%E6%94%AF%E8%AF%B4%E6%98%8E">返回主分支 ↩</a>
</div>

- [x] 服务发现
- [x] gRPC + protobuf
- [ ] API 网关

### 服务运行

> 直接通过 docker 启动：`docker-compose up -d`

+ 启用 服务发现：

    ```sh
    $ consul agent -dev
    ```

+ 启动 微服务模块

    ```sh
    user$ go run main.go
    order$ go run main.go
    ```

### 客户端访问

User 模块：`127.0.0.1:8081`
Order 模块：`127.0.0.1:8082`

### 端口占用

| 端口 | 用途 |
| ---- | ---- |
| 8081 | User Web 服务 |
| 8082 | Order Web 服务 |
| 18081 | User gRPC 服务 |
| 18082 | Order gRPC 服务 |

### 写在技术文档里的

由于微服务之间通过 rpc 调用远程服务，因此，每个微服务需要提供 http 服务给客户端的同时，需要额外提供 rpc 服务给其他微服务。

需要 API 网关的原因就是因为可以统一 Web 服务入口，而不需要每个微服务都提供一个 Web 服务入口。

### 二次开发


```go
// 需要到下载位置进行 `go build` 获取可执行文件
go get github.com/golang/protobuf
```

+ protoc 下载

https://github.com/protocolbuffers/protobuf/releases


+ protoc-gen-micro 下载

```sh
git clone https://github.com/gurufocus/protoc-gen-micro.git
cd protoc-gen-micro
go build
# protoc-gen-micro.exe
```

+ protoc-gen-go 下载

```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

```s
cd ./service/pb/
protoc --proto_path=. --micro_out=../ --go_out=../ userModel.proto
protoc --proto_path=. --micro_out=../ --go_out=../ userService.proto
```

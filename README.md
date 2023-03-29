## 微服务应用代码示例

<br>

<div align="right">
    <a href="https://github.com/fmw666/microservice-code-sample/tree/master#-%E5%88%86%E6%94%AF%E8%AF%B4%E6%98%8E">返回主分支 ↩</a>
</div>

- [ ] 服务发现（Consul）
- [ ] RPC + protobuf
- [ ] API 网关
- [ ] 消息队列（RabbitMQ）

```s
docker-compose up -d
```

```go
go get github.com/micro/go-micro/v2

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
protoc --proto_path=. --micro_out=../ --go_out=../ userModel.proto
protoc --proto_path=. --micro_out=../ --go_out=../ userService.proto
```

### 服务运行

+ 启用 服务发现：

    ```sh
    $ consul agent -dev -node fmw
    ```

+ 启动 微服务模块

    ```sh
    user$ go run main.go
    order$ go run main.go
    ```

### 服务访问

user 服务：<http://localhost:8081>

> swagger 文档：<http://localhost:8081/swagger/index.html>

order 服务：<http://localhost:8082>

> swagger 文档：<http://localhost:8082/swagger/index.html>

### 项目结构

```swift
.
├── order/                                   // order 服务
│  ├── api/                                  // 对外的接口，所有 http 请求的入口
│  │  ├── v1/                                // api 版本号
│  │  │  └── order.go                        // 订单接口
│  │  ├── v2/
│  │  └── ...
│  │
│  ├── config/                                // 配置文件包
│  │  ├── conf.ini                            // 配置文件
│  │  ├── config.go                           // 解析配置文件
│  │  └── section.go                          // 配置文件中的 section 定义为结构体
│  │
│  ├── docs/                                  // 由 swag 生成的文档
│  │  └── ...
│  │
│  ├── logs/                                  // 日志文件
│  │  └── ...
│  │
│  ├── middleware/                            // 中间件包
│  │  ├── cors.go                             // 跨域中间件
│  │  └── init.go                             // 错误处理中间件
│  │
│  │── models/                                // 模型包
│  │   ├── common.go                          // 公共模型、字段
│  │   ├── init.go                            // 初始化 DB
│  │   └── order.go                           // 订单模型
│  │
│  ├── pkg/                                   // 项目通用包
│  │   │── e/                                 // 错误处理包
│  │   │   │── code.go                        // 错误码
│  │   │   └── msg.go                         // 错误消息
│  │   │
│  │   │── logger/                            // 日志包
│  │   │   │── file.go                        // 日志文件
│  │   │   └── log.go                         // 日志记录器
│  │   │
│  │   └── utils/                             // 工具包
│  │       │── auth.go                        // 认证工具
│  │       └── response.go                    // 响应工具
│  │
│  ├── router/                                // 路由包
│  │   └── router.go                          // 路由初始化
│  │
│  ├── schema/                                // 请求/响应结构
│  │   ├── page.go                            // 分页请求/响应结构
│  │   └── user.go                            // user 请求/响应结构
│  │
│  ├── service/                               // 服务包
│  │   └── order.go                           // 订单服务
│  │
│  │── go.mod                                 // go module
│  └── main.go                                // 入口文件
│
├── user/                                     // user 服务
│  ├── api/                                   // 对外的接口，所有 http 请求的入口
│  │  ├── v1/                                 // api 版本号
│  │  │  └── user.go                          // 用户接口
│  │  ├── v2/
│  │  └── ...
│  │
│  ├── config/                                // 配置文件包
│  │  ├── conf.ini                            // 配置文件
│  │  ├── config.go                           // 解析配置文件
│  │  └── section.go                          // 配置文件中的 section 定义为结构体
│  │
│  ├── docs/                                  // 由 swag 生成的文档
│  │  └── ...
│  │
│  ├── logs/                                  // 日志文件
│  │  └── ...
│  │
│  ├── middleware/                            // 中间件包
│  │  ├── auth.go                             // 认证中间件
│  │  ├── cors.go                             // 跨域中间件
│  │  └── init.go                             // 错误处理中间件
│  │
│  │── models/                                // 模型包
│  │   ├── common.go                          // 公共模型、字段
│  │   ├── init.go                            // 初始化 DB
│  │   └── user.go                            // 用户模型
│  │
│  ├── pkg/                                   // 项目通用包
│  │   │── e/                                 // 错误处理包
│  │   │   │── code.go                        // 错误码
│  │   │   └── msg.go                         // 错误消息
│  │   │
│  │   │── logger/                            // 日志包
│  │   │   │── file.go                        // 日志文件
│  │   │   └── log.go                         // 日志记录器
│  │   │
│  │   └── utils/                             // 工具包
│  │       │── auth.go                        // 认证工具
│  │       └── response.go                    // 响应工具
│  │
│  ├── router/                                // 路由包
│  │   └── router.go                          // 路由初始化
│  │
│  ├── schema/                                // 请求/响应结构
│  │   ├── page.go                            // 分页请求/响应结构
│  │   └── user.go                            // user 请求/响应结构
│  │
│  ├── service/                               // 服务包
│  │   └── user.go                            // 用户服务
│  │
│  │── go.mod                                 // go module
│  └── main.go                                // 入口文件
│
└── manage.go                                // 服务管理文件
```

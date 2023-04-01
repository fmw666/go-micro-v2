## 微服务应用代码示例

<br>

<div align="right">
    <a href="https://github.com/fmw666/microservice-code-sample/tree/master#-%E5%88%86%E6%94%AF%E8%AF%B4%E6%98%8E">返回主分支 ↩</a>
</div>

- [x] 服务发现（Consul）
- [x] RPC + protobuf
- [ ] API 网关
- [ ] 消息队列（RabbitMQ）

### 项目部署

> 仅推荐 Docker 部署，其他部署方式请自行解决依赖问题

```s
docker-compose up -d
```

### 服务运行

> ❗ docker 启动后服务已主动开启

+ 启用 服务发现

    ```sh
    $ consul agent -dev
    ```

+ 启动 微服务模块

    *使用 manage.go 管理文件*

    ```sh
    code$ go run manage.go init
    code$ go run manage.go start
    ```

    *在特定服务中启动*

    ```sh
    code/user$ go run main.go
    code/order$ go run main.go
    ```

### 服务访问

+ 服务发现 服务：<http://localhost:8500>
    + 服务名：`consul`

+ user 服务：<http://localhost:8081>
    + swagger 文档：<http://localhost:8081/swagger/index.html>
    + 服务名：`rpcUserService`

+ order 服务：<http://localhost:8082>
    + swagger 文档：<http://localhost:8082/swagger/index.html>
    + 服务名：`rpcOrderService`

+ 端口占用

    | 端口 | 用途 |
    | ---- | ---- |
    | 8081 | User Web 服务 |
    | 8082 | Order Web 服务 |
    | 18081 | User gRPC 服务 |
    | 18082 | Order gRPC 服务 |

    *由于微服务之间通过 RPC 调用远程服务，每个微服务需要同时提供 HTTP 服务给客户端和额外提供 RPC 服务给其他微服务。为了统一 Web 服务入口，避免每个微服务都提供一个 Web 服务入口，需要使用 API 网关。*

### 二次开发

> 目前相关工具版本不太统一，依赖问题请自行解决

+ 依赖安装？？？

    ```go
    // 需要到下载位置进行 `go build` 获取可执行文件
    go get github.com/golang/protobuf
    ```

    + protoc 下载

        ```sh
        $ apt update
        $ apt install golang-goprotobuf-dev
        ```

    + protoc-gen-micro 下载

        *Windows 环境*

        ```sh
        git clone https://github.com/gurufocus/protoc-gen-micro.git
        cd protoc-gen-micro
        go build
        # protoc-gen-micro.exe
        ```

        *Linux 环境*

        ???

    + protoc-gen-go 下载

        ```sh
        $ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
        ```

+ 生成 pb 文件

    ```s
    $ cd ./service/pb/
    $ protoc --proto_path=. --micro_out=../ --go_out=../ userModel.proto
    $ protoc --proto_path=. --micro_out=../ --go_out=../ userService.proto
    ```

### 项目结构

+ 概览

    ```swift
    .
    ├── order/                                    // order 服务
    │  ├── ...                                    // 服务相关
    │  └── main.go                                // 入口文件
    │
    ├── user/                                     // user 服务
    │  ├── ...                                    // 服务相关
    │  └── main.go                                // 入口文件
    │
    └── manage.go                                 // 服务管理文件
    ```

+ 详情

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
    │  ├── core/                                  // 服务核心包
    │  │  ├── orderService.go                     // 对外提供的 order 服务
    │  │  └── pkg.go                              // 结构定义
    │  │
    │  ├── docs/                                  // 由 swag 生成的文档
    │  │  └── ...
    │  │
    │  ├── logs/                                  // 日志文件
    │  │  └── ...
    │  │
    │  ├── middleware/                            // 中间件包
    │  │  ├── cors.go                             // 跨域中间件
    │  │  ├── init.go                             // 错误处理中间件
    │  │  └── logger.go                           // 请求日志中间件
    │  │
    │  │── models/                                // 模型包
    │  │  ├── common.go                           // 公共模型、字段
    │  │  ├── init.go                             // 初始化 DB
    │  │  └── order.go                            // 订单模型
    │  │
    │  ├── pkg/                                   // 项目通用包
    │  │  │── e/                                  // 错误处理包
    │  │  │   │── code.go                         // 错误码
    │  │  │   └── msg.go                          // 错误消息
    │  │  │
    │  │  │── logger/                             // 日志包
    │  │  │   │── file.go                         // 日志文件
    │  │  │   └── log.go                          // 日志记录器
    │  │  │
    │  │  └── utils/                              // 工具包
    │  │      │── auth.go                         // 认证工具
    │  │      └── response.go                     // 响应工具
    │  │
    │  ├── router/                                // 路由包
    │  │  └── router.go                           // 路由初始化
    │  │
    │  ├── schema/                                // 请求/响应结构
    │  │  ├── codec.go                            // 微服务请求/响应结构
    │  │  ├── order.go                            // order 请求/响应结构
    │  │  └── page.go                             // 分页请求/响应结构
    │  │
    │  ├── service/                               // 服务包
    │  │  ├── pb/                                 // protobuf 文件
    │  │  │  │── orderModel.proto                 // order 模型
    │  │  │  │── orderService.proto               // order 服务
    │  │  │  │── userModel.proto                  // user 模型
    │  │  │  └── userService.proto                // user 服务
    │  │  │
    │  │  │── orderModel.pb.go                    // protoc 生成的 order 模型
    │  │  │── orderModel.pb.micro.go              // protoc 生成的基于 micro 的 order 模型
    │  │  │── orderService.pb.go                  // protoc 生成的 order 服务
    │  │  └── orderService.pb.micro.go            // protoc 生成的基于 micro 的 order 服务
    │  │
    │  ├── wrappers/                              // 微服务包装器
    │  │  └── orderWrapper.go                     // order 服务包装器
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
    │  ├── core/                                  // 服务核心包
    │  │  ├── pkg.go                              // 结构定义
    │  │  └── userService.go                      // 对外提供的 user 服务
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
    │  │  ├── init.go                             // 错误处理中间件
    │  │  └── logger.go                           // 请求日志中间件
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
    │  │  ├── codec.go                            // 微服务请求/响应结构
    │  │  ├── order.go                            // order 请求/响应结构
    │  │  ├── page.go                             // 分页请求/响应结构
    │  │  └── user.go                             // user 请求/响应结构
    │  │
    │  ├── service/                               // 服务包
    │  │  ├── pb/                                 // protobuf 文件
    │  │  │  │── orderModel.proto                 // order 模型
    │  │  │  │── orderService.proto               // order 服务
    │  │  │  │── userModel.proto                  // user 模型
    │  │  │  └── userService.proto                // user 服务
    │  │  │
    │  │  │── orderModel.pb.go                    // protoc 生成的 order 模型
    │  │  │── orderModel.pb.micro.go              // protoc 生成的基于 micro 的 order 模型
    │  │  │── orderService.pb.go                  // protoc 生成的 order 服务
    │  │  │── orderService.pb.micro.go            // protoc 生成的基于 micro 的 order 服务
    │  │  │── userModel.pb.go                     // protoc 生成的 user 模型
    │  │  │── userModel.pb.micro.go               // protoc 生成的基于 micro 的 user 模型
    │  │  │── userService.pb.go                   // protoc 生成的 user 服务
    │  │  └── userService.pb.micro.go             // protoc 生成的基于 micro 的 user 服务
    │  │
    │  ├── wrappers/                              // 微服务包装器
    │  │  │── orderWrapper.go                     // order 服务包装器
    │  │  └── userWrapper.go                      // user 服务包装器
    │  │
    │  │── go.mod                                 // go module
    │  └── main.go                                // 入口文件
    │
    └── manage.go                                 // 服务管理文件
    ```

## 微服务应用代码示例

<br>

<div align="right">
    <a href="https://github.com/fmw666/microservice-code-sample/tree/master#-%E5%88%86%E6%94%AF%E8%AF%B4%E6%98%8E">返回主分支 ↩</a>
</div>

- [x] 服务发现（Consul）
- [x] RPC + protobuf
- [x] API 网关
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

+ 启动 API网关模块

    ```sh
    code/api-gateway$ go run main.go
    ```

### 服务访问

+ 网关 服务：<http://localhost:8080>
    + swagger 文档：<http://localhost:8080/swagger/index.html>
    + 服务名：`httpService`

+ user 服务：<http://localhost:8081>
    + 服务名：`rpcUserService`

+ order 服务：<http://localhost:8082>
    + 服务名：`rpcOrderService`

+ 服务发现 服务：<http://localhost:8500>
    + 服务名：`consul`

+ 端口占用

    | 端口 | 用途 |
    | ---- | ---- |
    | 8080 | API 网关 服务 |
    | 8081 | User gRPC 服务 |
    | 8082 | Order gRPC 服务 |
    | 8500 | 服务发现 服务 |

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

        or

        curl -L -o /tmp/protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v3.19.1/protoc-3.19.1-linux-x86_64.zip && \
        unzip -d /tmp/protoc /tmp/protoc.zip && \
        mv /tmp/protoc/bin/protoc $GOPATH/bin
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

+ 处理 omitempty 属性

    ```s
    protoc-go-inject-tag -input="*.pb.go"
    ```

### 项目结构

+ 概览

    ```swift
    .
    ├── api-garteway/                             // 网关 服务
    │  ├── ...                                    // 网关服务相关
    │  └── main.go                                // 入口文件
    │
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
    ├── api-gateway/                              // 网关 服务
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
    │  ├── pkg/                                   // 项目通用包
    │  │  ├── logger/                             // 日志包
    │  │  │  ├── file.go                          // 日志文件
    │  │  │  └── log.go                           // 日志记录器
    │  │  │
    │  │  └── utils/                              // 工具包
    │  │     └── auth.go                          // 认证工具
    │  │
    │  ├── service/                               // 服务包
    │  │  ├── pb/                                 // protobuf 文件
    │  │  │  │── orderService.proto               // order 服务
    │  │  │  └── userService.proto                // user 服务
    │  │  │
    │  │  │── orderService.pb.go                  // protoc 生成的 order 服务
    │  │  │── orderService.pb.micro.go            // protoc 生成的基于 micro 的 order 服务
    │  │  │── userService.pb.go                   // protoc 生成的 user 服务
    │  │  └── userService.pb.micro.go             // protoc 生成的基于 micro 的 user 服务
    │  │
    │  ├── weblib/                                // web 服务库
    │  │  ├── handlers/                           // 对外的接口，http 请求的入口
    │  │  │  ├── order.go                         // 订单接口
    │  │  │  ├── pkg.go                           // 错误包装器
    │  │  │  └── user.go                          // 用户接口
    │  │  │
    │  │  ├── middleware/                         // 中间件包
    │  │  │  ├── auth.go                          // 认证中间件
    │  │  │  ├── cors.go                          // 跨域中间件
    │  │  │  ├── init.go                          // 错误处理中间件
    │  │  │  └── logger.go                        // 请求日志中间件
    │  │  │
    │  │  ├── schema/                             // 请求/响应结构
    │  │  │  ├── order.go                         // order 请求/响应结构
    │  │  │  └── user.go                          // user 请求/响应结构
    │  │  │
    │  │  └── router.go                           // 路由
    │  │
    │  ├── wrappers/                              // 微服务包装器
    │  │  ├── orderWrapper.go                     // order 服务包装器
    │  │  └── userWrapper.go                      // user 服务包装器
    │  │
    │  │── go.mod                                 // go module
    │  └── main.go                                // 入口文件
    │
    ├── order/                                    // order 服务
    │  ├── config/                                // 配置文件包
    │  │  ├── conf.ini                            // 配置文件
    │  │  ├── config.go                           // 解析配置文件
    │  │  └── section.go                          // 配置文件中的 section 定义为结构体
    │  │
    │  ├── core/                                  // 服务核心包
    │  │  ├── codec.go                            // 微服务请求/响应结构定义
    │  │  ├── orderService.go                     // 对外提供的 order 服务
    │  │  └── pkg.go                              // 结构定义
    │  │
    │  ├── logs/                                  // 日志文件
    │  │  └── ...
    │  │
    │  │── models/                                // 模型包
    │  │  ├── common.go                           // 公共模型、字段
    │  │  ├── init.go                             // 初始化 DB
    │  │  └── order.go                            // 订单模型
    │  │
    │  ├── pkg/                                   // 项目通用包
    │  │  ├── e/                                  // 错误处理包
    │  │  │   ├── code.go                         // 错误码
    │  │  │   └── msg.go                          // 错误消息
    │  │  │
    │  │  └── logger/                             // 日志包
    │  │      ├── file.go                         // 日志文件
    │  │      └── log.go                          // 日志记录器
    │  │
    │  ├── service/                               // 服务包
    │  │  ├── pb/                                 // protobuf 文件
    │  │  │  └── orderService.proto               // order 服务
    │  │  │
    │  │  │── orderService.pb.go                  // protoc 生成的 order 服务
    │  │  └── orderService.pb.micro.go            // protoc 生成的基于 micro 的 order 服务
    │  │
    │  │── go.mod                                 // go module
    │  └── main.go                                // 入口文件
    │
    ├── user/                                     // user 服务
    │  ├── config/                                // 配置文件包
    │  │  ├── conf.ini                            // 配置文件
    │  │  ├── config.go                           // 解析配置文件
    │  │  └── section.go                          // 配置文件中的 section 定义为结构体
    │  │
    │  ├── core/                                  // 服务核心包
    │  │  ├── codec.go                            // 微服务请求/响应结构定义
    │  │  ├── pkg.go                              // 结构定义
    │  │  └── userService.go                      // 对外提供的 user 服务
    │  │
    │  ├── logs/                                  // 日志文件
    │  │  └── ...
    │  │
    │  │── models/                                // 模型包
    │  │   ├── common.go                          // 公共模型、字段
    │  │   ├── init.go                            // 初始化 DB
    │  │   └── user.go                            // 用户模型
    │  │
    │  ├── pkg/                                   // 项目通用包
    │  │  ├── e/                                  // 错误处理包
    │  │  │  ├── code.go                          // 错误码
    │  │  │  └── msg.go                           // 错误消息
    │  │  │
    │  │  ├── logger/                             // 日志包
    │  │  │  ├── file.go                          // 日志文件
    │  │  │  └── log.go                           // 日志记录器
    │  │  │
    │  │  └── utils/                              // 工具包
    │  │     └── auth.go                          // 认证工具
    │  │
    │  ├── service/                               // 服务包
    │  │  ├── pb/                                 // protobuf 文件
    │  │  │  └── userService.proto                // user 服务
    │  │  │
    │  │  │── userService.pb.go                   // protoc 生成的 user 服务
    │  │  └── userService.pb.micro.go             // protoc 生成的基于 micro 的 user 服务
    │  │
    │  │── go.mod                                 // go module
    │  └── main.go                                // 入口文件
    │
    └── manage.go                                 // 服务管理文件
    ```

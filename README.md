## 单体应用代码示例

### Docker 部署

```s
docker-compose up -d
```

### Docker 中运行

初始化数据库、swag、mod

source script init



### 本地运行

### 服务运行


+ 启动 API网关模块

    ```sh
    code$ go run main.go
    ```

### 项目结构

```swift
.
├── api/                                   // 对外的接口，所有 http 请求的入口
│  ├── v1/                                 // api 版本号
│  │  ├── order.go                         // 订单接口
│  │  └── user.go                          // 用户接口
│  ├── v2/
│  └── ...
│
├── config/                                // 配置文件包
│  ├── conf.ini                            // 配置文件
│  ├── config.go                           // 解析配置文件
│  └── section.go                          // 配置文件中的 section 定义为结构体
│
├── docs/                                  // 由 swag 生成的文档
│  └── ...
│
├── logs/                                  // 日志文件
│  └── ...
│
├── middleware/                            // 中间件包
│  ├── auth.go                             // 认证中间件
│  ├── cors.go                             // 跨域中间件
│  └── init.go                             // 错误处理中间件
│
│── models/                                // 模型包
│   ├── common.go                          // 公共模型、字段
│   ├── init.go                            // 初始化 DB
│   ├── order.go                           // 订单模型
│   └── user.go                            // 用户模型
│
├── pkg/                                   // 项目通用包
│   │── e/                                 // 错误处理包
│   │   │── code.go                        // 错误码
│   │   └── msg.go                         // 错误消息
│   │
│   │── logger/                            // 日志包
│   │   │── file.go                        // 日志文件
│   │   └── log.go                         // 日志记录器
│   │
│   └── utils/                             // 工具包
│       └── auth.go                        // 认证工具
│
├── router/                                // 路由包
│   ├── order.go                           // 订单路由
│   ├── ping.go                            // ping 测试路由
│   ├── router.go                          // 路由初始化
│   └── user.go                            // 用户路由
│
├── schema/                                // 请求/响应结构
│   ├── order/                             // order 请求/响应结构
│   │   │── request.go
│   │   └── response.go
│   │
│   ├── user/                              // user 请求/响应结构
│   │   │── request.go
│   │   └── response.go
│   │
│   ├── base.go                            // 基本请求/响应结构
│   └── page.go                            // 分页请求/响应结构
│
├── service/                               // 服务包
│   ├── order.go                           // 订单服务
│   └── user.go                            // 用户服务
│
│── go.mod                                 // go module
└── main.go                                // 入口文件
```

## 微服务应用代码示例

### 🚀 项目介绍

+ 基础介绍

    | 项目地址 | URL 前缀 | 数据库 |
    | -------- | -------- | -------- |
    | 见具体分支 | /api/v1 | MySQL<br>库名：mall<br>表名：user, order |

+ 项目技术

    | 编程语言 | web 框架 | ORM | 认证 | 日志 | 微服务框架 |
    | ------- | ---- | ---- | ---- | ---- | ---- |
    | Go 1.18 | gin | gorm | JWT | logrus | go-micro v2 |

+ User 服务

    + 端口：8081

    + 模型：User

        > 表名：user

        | 字段 | 类型 | 备注 |
        | ---- | ---- | ---- |
        | id | int | 主键 |
        | created_at | datetime | 创建时间 |
        | updated_at | datetime | 更新时间 |
        | deleted_at | datetime | 删除时间 |
        | username | string | 用户名 |
        | password | string | 密码 |

    + 接口

        | 接口名 | 请求方式 | 请求路径 | 请求参数 | 返回值 |
        | ---- | ---- | ---- | ---- | ---- |
        | 注册 | POST | /user/register | username, password | 注册成功，返回用户信息 |
        | 登录 | POST | /user/login | username, password | 登录成功，返回用户信息 |
        | 创建订单 | POST | /users/{:user_id}/orders | user_id, order_id | 创建成功，返回订单信息 |
        | 查询订单 | GET | /users/{:user_id}/orders | user_id, order_id | 查询成功，返回订单信息 |

+ Order 服务

    + 端口：8082

    + 模型：Order

        > 表名：order

        | 字段 | 类型 | 备注 |
        | ---- | ---- | ---- |
        | id | int | 主键 |
        | created_at | datetime | 创建时间 |
        | updated_at | datetime | 更新时间 |
        | deleted_at | datetime | 删除时间 |
        | name | string | 订单名称 |
        | user_id | int | 用户id |


    + 接口（需要 JWT 认证）

        | 接口名 | 请求方式 | 请求路径 | 请求参数 | 返回值 |
        | ---- | ---- | ---- | ---- | ---- |
        | 创建订单 | POST | /orders | name, user_id | 创建成功，返回订单信息 |
        | 获取订单列表 | GET | /orders | user_id | 获取成功，返回订单列表 |

### 🎈 分支说明

> swagger-ui 应用在 API 入口

| 分支名称 | 应用端口 | API 入口 | 分支描述 |
| ------- | -------- | ------- | ------- |
| master | - | - | 描述文档 |
| monolithic-app | 8080 | 8080 | 单体应用代码 |
| microservice-app | User 服务：8081<br>Order 服务：8082 | 8081, 8082 | 微服务应用代码 |
| microservice-app-with-gateway | API 网关服务：8080<br>User 服务：8081<br>Order 服务：8082 | 8080 | 微服务应用代码，同时支持网关 |

<div align="center">
  <h1><code>微服务应用代码示例</code></h1>

  <p>
    <strong>目录导航：
    <a href="#-项目介绍">🚀 项目介绍</a></strong>
    <span>&ensp;|&ensp;</span>
    <a href="#-分支说明">🎈 分支说明</a></strong>
  </p>

  <strong>基于：<a href="https://bytecodealliance.org/">《一篇文章让你了解微服务架构设计》</a></strong>

  <p>
    <a href="https://github.com/bytecodealliance/wasmtime-go/actions?query=workflow%3ACI">
      <img src="https://github.com/bytecodealliance/wasmtime-go/workflows/CI/badge.svg" alt="CI status"/>
    </a>
    <a href="https://pkg.go.dev/github.com/bytecodealliance/wasmtime-go">
      <img src="https://godoc.org/github.com/bytecodealliance/wasmtime-go?status.svg" alt="Documentation"/>
    </a>
    <a href="https://bytecodealliance.github.io/wasmtime-go/coverage.html">
      <img src="https://img.shields.io/badge/coverage-main-green" alt="Code Coverage"/>
    </a>
  </p>

</div>

## 🚀 项目介绍

+ **[项目概要](#no-reply)**

    基于商场类项目中的订单模块。该模块中包含两个服务：一是用户服务，用户可以进行注册登录、登录后可以进行订单的查询、新增操作；二是订单服务，可以进行订单的查询、新增操作。

+ **[项目技术](#no-reply)**

    | 编程语言 | web 框架 | 数据库 | ORM | 认证 | 日志 | 微服务框架 |
    | ------- | ---- | ---- | ---- | ---- | ---- | ---- |
    | Go 1.18 | gin | MySQL | gorm | JWT | logrus | go-micro v2 |

+ User 服务

    + 微服务应用端口：8081

    + 微服务应用名称：userRpcService


        <li>
        <details>
        <summary>模型：User</summary>
        <blockquote>
        <p dir="auto">表名：user</p>
        </blockquote>
        <table>
        <thead>
        <tr>
        <th>字段</th>
        <th>类型</th>
        <th>备注</th>
        </tr>
        </thead>
        <tbody>
        <tr>
        <td>id</td>
        <td>int</td>
        <td>主键</td>
        </tr>
        <tr>
        <td>created_at</td>
        <td>datetime</td>
        <td>创建时间</td>
        </tr>
        <tr>
        <td>updated_at</td>
        <td>datetime</td>
        <td>更新时间</td>
        </tr>
        <tr>
        <td>deleted_at</td>
        <td>datetime</td>
        <td>删除时间</td>
        </tr>
        <tr>
        <td>username</td>
        <td>string</td>
        <td>用户名</td>
        </tr>
        <tr>
        <td>password</td>
        <td>string</td>
        <td>密码</td>
        </tr>
        </tbody>
        </table>
        </details>
        </li>

    + 接口

        | 接口名 | 请求方式 | 请求路径 | 请求参数 | 返回值 |
        | ---- | ---- | ---- | ---- | ---- |
        | 注册 | POST | /user/register | username, password | 注册成功，返回用户信息 |
        | 登录 | POST | /user/login | username, password | 登录成功，返回用户信息 |
        | 创建订单 | POST | /users/{:user_id}/orders | user_id, order_id | 创建成功，返回订单信息 |
        | 查询订单 | GET | /users/{:user_id}/orders | user_id, order_id | 查询成功，返回订单信息 |

+ Order 服务

    + 微服务应用端口：8082

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

## 🎈 分支说明

> swagger-ui 应用在 API 入口

| 分支名称 | 应用端口 | API 入口 | 分支描述 |
| ------- | -------- | ------- | ------- |
| master | - | - | 描述文档 |
| monolithic-app | 8080 | 8080 | 单体应用代码 |
| microservice-app | User 服务：8081<br>Order 服务：8082 | 8081, 8082 | 微服务应用代码 |
| microservice-app-with-gateway | API 网关服务：8080<br>User 服务：8081<br>Order 服务：8082 | 8080 | 微服务应用代码，同时支持网关 |

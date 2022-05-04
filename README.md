<div align="center">
  <h1>💬 微服务应用代码示例</h1>

  <div align="right">
  <p>
    <a href="https://golang.google.cn/">
      <img src="https://img.shields.io/badge/-%20
      1.18-007D9C?logo=go&logoColor=white&style=flat&logoWidth=16" alt="language"/>
    </a>
    <a href="#">
      <img src="https://badgen.net/github/branches/fmw666/microservice-code-sample" alt="branches"/>
    </a>
    <a href="https://github.com/fmw666/microservice-code-sample/blob/master/LICENSE">
      <img src="https://badgen.net/github/license/fmw666/microservice-code-sample" alt="license"/>
    </a>
  </p>

  <strong>基于：<a href="https://bytecodealliance.org/">《一篇文章让你了解微服务架构设计》</a></strong>

  <p>
    <strong>目录导航：
    <li><a href="#-项目介绍">🚀 项目介绍</a></li>
    <li><a href="#-分支说明">🎈 分支说明</a></li>
  </p>

  </div>

</div>

## 🚀 项目介绍

+ **[项目概要](#no-reply)**

    基于商场类项目中的订单模块。该模块中包含两个服务：一是用户服务，用户可以进行注册登录、登录后可以进行订单的查询、新增操作；二是订单服务，可以进行订单的查询、新增操作。

+ **[项目技术](#no-reply)**

    | 编程语言 | web 框架 | 数据库 | ORM | 认证 | 日志 | 微服务框架 |
    | ------- | ---- | ---- | ---- | ---- | ---- | ---- |
    | Go 1.18 | gin | MySQL | gorm | JWT | logrus | go-micro v2 |

+ **[User 服务](#no-reply)**

    + 微服务应用端口：`8081`

    + 微服务应用名称：`userRpcService`

        <li>
        <details>
        <p dir="auto"><summary>模型：<code>User</code></summary></p>
        <blockquote>
        <p dir="auto">表名：<code>user</code></p>
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

        <li>
        <details>
        <p dir="auto"><summary>接口</summary></p>
        <table>
        <thead>
        <tr>
        <th>接口名</th>
        <th>请求方式</th>
        <th>请求路径</th>
        <th>请求参数</th>
        <th>返回值</th>
        </tr>
        </thead>
        <tbody>
        <tr>
        <td>注册</td>
        <td>POST</td>
        <td>/user/register</td>
        <td>username, password</td>
        <td>注册成功，返回用户信息</td>
        </tr>
        <tr>
        <td>登录</td>
        <td>POST</td>
        <td>/user/login</td>
        <td>username, password</td>
        <td>登录成功，返回用户信息</td>
        </tr>
        <tr>
        <td>创建订单</td>
        <td>POST</td>
        <td>/users/{:user_id}/orders</td>
        <td>user_id, order_id</td>
        <td>创建成功，返回订单信息</td>
        </tr>
        <tr>
        <td>查询订单</td>
        <td>GET</td>
        <td>/users/{:user_id}/orders</td>
        <td>user_id, order_id</td>
        <td>查询成功，返回订单信息</td>
        </tr>
        </tbody>
        </table>
        </details>
        </li>

+ **[Order 服务](#no-reply)**

    + 微服务应用端口：`8082`

    + 微服务应用名称：`orderRpcService`

        <li>
        <details>
        <p dir="auto"><summary>模型：<code>Order</code></summary></p>
        <blockquote>
        <p dir="auto">表名：<code>order</code></p>
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
        <td>name</td>
        <td>string</td>
        <td>订单名称</td>
        </tr>
        <tr>
        <td>user_id</td>
        <td>int</td>
        <td>用户id</td>
        </tr>
        </tbody>
        </table>
        </details>
        </li>

        <li>
        <details>
        <p dir="auto"><summary>接口（需要 JWT 认证）</summary></p>
        <table>
        <thead>
        <tr>
        <th>接口名</th>
        <th>请求方式</th>
        <th>请求路径</th>
        <th>请求参数</th>
        <th>返回值</th>
        </tr>
        </thead>
        <tbody>
        <tr>
        <td>创建订单</td>
        <td>POST</td>
        <td>/orders</td>
        <td>name, user_id</td>
        <td>创建成功，返回订单信息</td>
        </tr>
        <tr>
        <td>获取订单列表</td>
        <td>GET</td>
        <td>/orders</td>
        <td>user_id</td>
        <td>获取成功，返回订单列表</td>
        </tr>
        </tbody>
        </table>
        </details>
        </li>

<br>

## 🎈 分支说明

> swagger-ui 应用在 API 入口

| 分支名称 | 应用端口 | API 入口 | 分支描述 |
| :------ | :------- | :------ | :------ |
| [master](https://github.com/fmw666/microservice-example/tree/master) | - | - | 描述文档 |
| [monolithic-app](https://github.com/fmw666/microservice-example/tree/monolithic-app) | 8080 | 8080 | 单体应用代码 |
| [microservice-app](https://github.com/fmw666/microservice-example/tree/microservice-app) | User 服务：8081<br>Order 服务：8082 | 8081, 8082 | 微服务应用代码 |
| [microservice-app-with-gateway](https://github.com/fmw666/microservice-example/tree/microservice-app-with-gateway) | API 网关服务：8080<br>User 服务：8081<br>Order 服务：8082 | 8080 | 微服务应用代码<br>使用网关作为入口 |

## 微服务应用代码示例

<br>

<div align="right">
    <a href="https://github.com/fmw666/microservice-code-sample/tree/master#-%E5%88%86%E6%94%AF%E8%AF%B4%E6%98%8E">返回主分支 ↩</a>
</div>

- [x] 服务发现
- [x] gRPC + protobuf
- [ ] API 网关

### Docker 运行

```s
docker-compose up -d
```


### 服务运行

+ 启用 服务发现：

    ```sh
    $ consul agent -dev
    ```

+ 启动 微服务模块

    ```sh
    user$ go run main.go
    order$ go run main.go
    ```

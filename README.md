## 微服务应用代码示例

<br>

- [x] 服务发现
- [ ] RPC + protobuf
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

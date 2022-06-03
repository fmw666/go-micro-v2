
docker 运行 rabbitmq

```sh
# 下载 rabbitmq management 镜像
docker pull rabbitmq:3-management

# 运行 rabbitmq 容器
docker run --rm -it --hostname my-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management

# 网页端入口
http://localhost:15672
(默认用户名: guest, 密码: guest)
```
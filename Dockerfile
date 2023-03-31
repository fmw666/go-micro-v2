FROM golang:1.18

ENV GOPROXY https://goproxy.cn,direct

# 安装必要的软件包和依赖包
USER root
RUN sed -i 's/deb.debian.org/mirrors.tuna.tsinghua.edu.cn/' /etc/apt/sources.list && \
    sed -i 's/security.debian.org/mirrors.tuna.tsinghua.edu.cn/' /etc/apt/sources.list && \
    sed -i 's/security-cdn.debian.org/mirrors.tuna.tsinghua.edu.cn/' /etc/apt/sources.list && \
    apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y --no-install-recommends \
    curl \
    zip \
    unzip \
    git \
    vim \
    screen \
    gnupg \
    apt-transport-https

# $GOPATH/bin 添加到环境变量中
ENV PATH $GOPATH/bin:$PATH

# 安装 swag
USER root
RUN apt-get install -y debian-keyring debian-archive-keyring apt-transport-https
RUN curl -1sLf 'https://dl.cloudsmith.io/public/go-swagger/go-swagger/gpg.2F8CB673971B5C9E.key' | gpg --dearmor -o /usr/share/keyrings/go-swagger-go-swagger-archive-keyring.gpg
RUN apt-get update
RUN go install github.com/swaggo/swag/cmd/swag@latest

# 清理垃圾
USER root
RUN apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* && \
    rm /var/log/lastlog /var/log/faillog

# 设置工作目录
WORKDIR /usr/src/code
CMD [ "go", "run", "manage.go", "init", "&&", "go", "run", "manage.go", "start" ]

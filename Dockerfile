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
RUN go install github.com/swaggo/swag/cmd/swag@latest

# 清理垃圾
USER root
RUN apt-get clean && \
    rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/* && \
    rm /var/log/lastlog /var/log/faillog

# 设置工作目录
WORKDIR /usr/src/code
# CMD [ "go", "run", "manage.go", "start" ]
# ENTRYPOINT [ "go", "run", "manage.go", "start" ]

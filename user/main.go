package main

import (
	"user/core"
	"user/service"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

func main() {
	// etcd 注册件
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	// 获取一个微服务的实例
	microService := micro.NewService(
		micro.Name("rpcUserService"),
		micro.Address("127.0.0.1:8082"),
		micro.Registry(etcdReg),
	)
	// 初始化
	microService.Init()
	// 注册微服务
	service.RegisterUserServiceHandler(microService.Server(), new(core.UserService))
	// 启动微服务
	microService.Run()
}

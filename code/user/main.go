package main

import (
	"user/core"
	"user/models"
	"user/service"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	// 初始化数据库
	models.Migrate()

	// consul 注册件
	consulReg := consul.NewRegistry(
		registry.Addrs(":8500"),
	)
	// 获取一个微服务的实例
	microService := micro.NewService(
		micro.Name("rpcUserService"),
		micro.Address("127.0.0.1:8081"),
		micro.Registry(consulReg),
	)
	// 服务注册
	service.RegisterUserServiceHandler(microService.Server(), new(core.UserService))
	// 启动微服务
	microService.Run()
}

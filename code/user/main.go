package main

import (
	"user/core"
	"user/models"
	"user/service"

	"github.com/aiscrm/go-micro/v2"
	"github.com/aiscrm/go-micro/v2/registry"
	"github.com/aiscrm/go-micro/v2/registry/consul"
)

func main() {
	// 初始化数据库
	models.Migrate()

	// consul 注册件
	consulReg := consul.NewRegistry(
		registry.Addrs("172.27.128.1:8500"),
	)
	// 获取一个微服务的实例
	microService := micro.NewService(
		micro.Name("rpcUserService"),
		micro.Address("0.0.0.0:8081"),
		micro.Registry(consulReg),
	)
	// 服务注册
	service.RegisterUserServiceHandler(microService.Server(), new(core.UserService))
	// 启动微服务
	microService.Run()
}

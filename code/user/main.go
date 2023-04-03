package main

import (
	"user/config"
	"user/core"
	"user/models"
	"user/pkg/logger"
	"user/service"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

func main() {
	logger.Info("user service start...")

	// 初始化数据库
	models.Migrate()

	// consul 注册件
	consulReg := consul.NewRegistry(
		registry.Addrs(config.ConsulSetting.Host + ":" + config.ConsulSetting.Port),
	)
	// 获取一个微服务的实例
	microService := micro.NewService(
		micro.Name(config.ServerSetting.MicroServiceName),
		micro.Address(config.ServerSetting.Host+":"+config.ServerSetting.RpcPort),
		micro.Registry(consulReg),
	)
	// 服务注册
	service.RegisterUserServiceHandler(microService.Server(), new(core.UserService))

	// 启动微服务
	microService.Run()
}

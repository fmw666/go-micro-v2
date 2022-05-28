package main

import (
	"order/config"
	"order/core"
	"order/models"
	"order/router"
	"order/service"
	"order/wrappers"
	"time"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

// 初始化 user、order 服务
func init_services(consulReg registry.Registry) map[string]any {
	userMicroService := micro.NewService(
		micro.Name("userService.client"),
		micro.WrapClient(wrappers.NewUserWrapper),
	)
	userService := service.NewUserService("rpcUserService", userMicroService.Client())

	orderMicroService := micro.NewService(
		micro.Name("orderService.client"),
		micro.WrapClient(wrappers.NewOrderWrapper),
	)
	orderService := service.NewOrderService("rpcOrderService", orderMicroService.Client())

	return map[string]any{
		"userService":  userService,
		"orderService": orderService,
	}
}

// 初始化 User 微服务
func init_microservice(consulReg registry.Registry) micro.Service {
	return micro.NewService(
		micro.Name(config.ServerSetting.MicroServiceName),
		micro.Address(config.ServerSetting.Host+":"+"18082"),
		// micro.Address(config.ServerSetting.Host+":"+config.ServerSetting.Port),
		micro.Registry(consulReg),
		// 设置注册服务过期时间
		micro.RegisterTTL(time.Second*30),
		// 设置间隔多久再次注册服务
		micro.RegisterInterval(time.Second*20),
	)
}

// @title Order API
// @version 1.0
// @description Order micro service.
// @host localhost:8082
// @BasePath /api/v1
func main() {
	// 初始化数据库
	models.Migrate()

	// consul 注册件
	consulReg := consul.NewRegistry(registry.Addrs(":8500"))

	// 初始化服务
	services := init_services(consulReg)

	// gin Router 路由引擎
	ginRouter := router.Router(services)

	// 获取 User 微服务的实例
	microService := init_microservice(consulReg)
	// 服务注册
	service.RegisterUserServiceHandler(microService.Server(), new(core.UserService))
	// 启动微服务
	go microService.Run()

	// 启动 web 服务
	ginRouter.Run(config.ServerSetting.Host + ":" + config.ServerSetting.Port)
}

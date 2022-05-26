package main

import (
	"time"
	"user/config"
	"user/core"
	"user/models"
	"user/pkg/utils/consul"
	"user/router"
	"user/service"
	"user/wrappers"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
)

func init_services(consulReg registry.Registry) map[string]any {
	// 初始化 user、order 服务
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

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @host localhost:8081
// @BasePath /api/v1
// @securityDefinitions.basic BasicAuth
// @in header
// @name Authorization
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// 初始化数据库
	models.Migrate()

	// consul 注册件
	consulReg := consul.ConsulReg

	// 初始化服务
	services := init_services(consulReg)

	// gin Router 路由引擎
	ginRouter := router.Router(services)

	// 获取一个微服务的实例
	microService := micro.NewService(
		micro.Name(config.ServerSetting.MicroServiceName),
		micro.Address(config.ServerSetting.Host+":"+"18081"),
		// micro.Address(config.ServerSetting.Host+":"+config.ServerSetting.Port),
		micro.Registry(consulReg),
		// 设置注册服务过期时间
		micro.RegisterTTL(time.Second*30),
		// 设置间隔多久再次注册服务
		micro.RegisterInterval(time.Second*20),
	)
	// 服务注册
	service.RegisterUserServiceHandler(microService.Server(), new(core.UserService))
	// 启动微服务
	go microService.Run()

	// 启动 web 服务
	ginRouter.Run(config.ServerSetting.Host + ":" + config.ServerSetting.Port)
}

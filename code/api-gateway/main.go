package main

import (
	"api-gateway/service"
	"api-gateway/weblib"
	"api-gateway/wrappers"
	"time"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/web"
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

// @title Swagger API-Gateway API 入口
// @version 1.0
// @description 网关 API 入口.
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.basic BasicAuth
// @in header
// @name Authorization
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// 初始化 consul 注册件
	consulReg := consul.NewRegistry(
		// registry.Addrs("172.27.128.1:8500"),
		registry.Addrs("127.0.0.1:8500"),
	)

	// 初始化服务
	services := init_services(consulReg)

	// 创建微服务实例，使用 gin 暴露 http 接口并注册到 consul
	microService := web.NewService(
		web.Name("httpService"),
		web.Address("127.0.0.1:8080"),
		// 将服务调用实例使用 gin 处理
		web.Handler(weblib.NewRouter(services)),
		web.Registry(consulReg),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
		web.Metadata(map[string]string{"protocol": "http"}),
	)

	// 接收命令行参数
	microService.Init()
	microService.Run()
}

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

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
func main() {
	// 初始化 consul 注册件
	consulReg := consul.NewRegistry(
		// registry.Addrs("172.27.128.1:8500"),
		registry.Addrs("127.0.0.1:8500"),
	)

	// 初始化 user 服务
	userMicroService := micro.NewService(
		micro.Name("userService.client"),
		micro.WrapClient(wrappers.NewUserWrapper),
	)
	// user 服务调用实例
	userService := service.NewUserService("rpcUserService", userMicroService.Client())

	// 初始化 order 服务
	orderMicroService := micro.NewService(
		micro.Name("orderService.client"),
		micro.WrapClient(wrappers.NewOrderWrapper),
	)
	// order 服务调用实例
	orderService := service.NewOrderService("rpcOrderService", orderMicroService.Client())

	// 创建微服务实例，使用 gin 暴露 http 接口并注册到 consul
	microService := web.NewService(
		web.Name("httpService"),
		web.Address("127.0.0.1:8080"),
		// 将服务调用实例使用 gin 处理
		web.Handler(weblib.NewRouter(userService, orderService)),
		web.Registry(consulReg),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
		web.Metadata(map[string]string{"protocol": "http"}),
	)
	// 接收命令行参数
	microService.Init()
	microService.Run()
}

package main

import (
	"api-gateway/service"
	"api-gateway/weblib"
	"api-gateway/wrappers"
	"time"

	"github.com/aiscrm/go-micro/v2"
	"github.com/aiscrm/go-micro/v2/registry"
	"github.com/aiscrm/go-micro/v2/registry/consul"
	"github.com/aiscrm/go-micro/v2/web"
)

func main() {
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
	// 用户
	userMicroService := micro.NewService(
		micro.Name("userService.client"),
		micro.WrapClient(wrappers.NewUserWrapper),
	)
	// 用户服务调用实例
	userService := service.NewUserService("userService", userMicroService.Client())

	// 创建微服务实例，使用 gin 暴露 http 接口并注册到 consul
	microService := web.NewService(
		web.Name("httpService"),
		web.Address("127.0.0.1:8080"),
		//将服务调用实例使用gin处理
		// web.Handler(weblib.NewRouter(userService, taskService)),
		web.Handler(weblib.NewRouter(userService)),
		web.Registry(consulReg),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
		web.Metadata(map[string]string{"protocol": "http"}),
	)
	//接收命令行参数
	microService.Init()
	microService.Run()
}

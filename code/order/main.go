package main

import (
	"order/core"
	"order/models"
	"order/service"

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
		micro.Name("rpcOrderService"),
		// web.RegisterTTL(30 * time.Second), // 设置注册服务的过期时间
		// web.RegisterInterval(20 * time.Second), // 设置间隔多久再次注册服务
		micro.Address("127.0.0.1:8082"),
		micro.Registry(consulReg),
	)
	// 服务注册
	service.RegisterOrderServiceHandler(microService.Server(), new(core.OrderService))

	// 启动微服务
	microService.Run()
}

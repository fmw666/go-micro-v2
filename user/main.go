package main

import (
	"os/exec"
	"user/core"
	_ "user/models"
	"user/service"

	"github.com/aiscrm/go-micro/v2"
	"github.com/aiscrm/go-micro/v2/registry"
	"github.com/aiscrm/go-micro/v2/registry/consul"
)

func init() {
	cmd := exec.Command("swag", "init")
	cmd.Run()
	// fmt.Println("swag init success")
}

func main() {
	// consul 注册件
	consulReg := consul.NewRegistry(
		registry.Addrs("127.0.0.1:8500"),
	)
	// 获取一个微服务的实例
	microService := micro.NewService(
		micro.Name("rpcUserService"),
		// web.RegisterTTL(30 * time.Second), // 设置注册服务的过期时间
		// web.RegisterInterval(20 * time.Second), // 设置间隔多久再次注册服务
		micro.Address("127.0.0.1:8081"),
		micro.Registry(consulReg),
	)
	// 服务注册
	service.RegisterUserServiceHandler(microService.Server(), new(core.UserService))
	// 启动微服务
	microService.Run()
}

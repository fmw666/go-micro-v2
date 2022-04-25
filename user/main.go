package main

import (
	"os/exec"
	"user/router"

	"github.com/aiscrm/go-micro/v2/registry"
	"github.com/aiscrm/go-micro/v2/registry/consul"
	"github.com/aiscrm/go-micro/v2/web"
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
	microService := web.NewService(
		web.Name("userService"),
		// web.RegisterTTL(30 * time.Second), // 设置注册服务的过期时间
		// web.RegisterInterval(20 * time.Second), // 设置间隔多久再次注册服务
		web.Address("127.0.0.1:8081"),
		web.Handler(router.Router()),
		web.Registry(consulReg),
	)
	// 启动微服务
	microService.Run()
}

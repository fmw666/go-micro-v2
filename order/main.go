package main

import (
	"bytes"
	"fmt"
	"net/http"
	"order/router"
	"os/exec"
	"time"

	"github.com/aiscrm/go-micro/v2/client/selector"
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
		web.Name("orderService"),
		// web.RegisterTTL(30 * time.Second), // 设置注册服务的过期时间
		// web.RegisterInterval(20 * time.Second), // 设置间隔多久再次注册服务
		web.Address("127.0.0.1:8082"),
		web.Handler(router.Router()),
		web.Registry(consulReg),
	)

	// 服务发现
	hostAddress := GetServiceAddr("userService", consulReg)
	if len(hostAddress) > 0 {
		url := "http://" + hostAddress + "/api/v1/users"
		response, _ := http.Post(url, "application/json;charset=utf-8", bytes.NewBuffer([]byte("{\"name\":\"test\",\"age\":18}")))

		fmt.Printf("发现服务，response = %v\n", response)
	}

	// 启动微服务
	microService.Run()
}

// 服务发现
func GetServiceAddr(serverName string, reg registry.Registry) (address string) {
	var retryTimes int
	for {
		servers, err := reg.GetService(serverName)
		fmt.Println("servers = ", servers)
		if err != nil {
			fmt.Println("GetServiceAddr err = ", err)
		}
		var services []*registry.Service
		for _, value := range servers {
			fmt.Println("value = ", value.Version)
			services = append(services, value)
		}
		// 获取其中一个服务的消息
		next := selector.RoundRobin(services)
		if node, err := next(); err == nil {
			address = node.Address
		}
		if len(address) > 0 {
			return
		}

		// 重试次数++
		retryTimes++
		time.Sleep(time.Second)
		// 重试 5 次，返回空
		if retryTimes >= 5 {
			return
		}
	}
}

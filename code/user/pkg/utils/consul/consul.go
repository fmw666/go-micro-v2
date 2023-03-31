package consul

import (
	"errors"
	"fmt"
	"time"
	"user/config"

	"github.com/micro/go-micro/v2/client/selector"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
)

// consul 注册件
var ConsulReg = initConsulReg()

func initConsulReg() registry.Registry {
	return consul.NewRegistry(
		registry.Addrs(config.ConsulSetting.Host + ":" + config.ConsulSetting.Port),
	)
}

// 服务发现
func GetServiceAddr(serverName string) (string, error) {
	var retryTimes int
	for {
		servers, err := ConsulReg.GetService(serverName)
		fmt.Println(servers)
		if err != nil {
			fmt.Println(err.Error())
		}
		var services []*registry.Service
		for _, value := range servers {
			fmt.Println(value.Name, ":", value.Version)
			services = append(services, value)
		}
		// 获取其中一个服务的信息
		var address string = ""
		next := selector.RoundRobin(services)
		if node, err := next(); err == nil {
			address = node.Address
		}
		if address != "" {
			return address, nil
		}

		// 重试次数++
		retryTimes++
		time.Sleep(1 * time.Second)
		// 重试5次 返回空
		if retryTimes >= 5 {
			return "", errors.New("get service failed")
		}
	}
}

package config

import (
	"fmt"
	"os"

	"github.com/go-ini/ini"
)

var cfg *ini.File

func init() {
	var err error
	cfg, err = ini.Load("config/conf.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
	}

	mapTo("server", ServerSetting)
	mapTo("consul", ConsulSetting)
	mapTo("service", ServiceSetting)

	// 从环境变量中读取 consul 配置
	if os.Getenv("CONSUL_HOST") != "" {
		ConsulSetting.Host = os.Getenv("CONSUL_HOST")
	}
	if os.Getenv("CONSUL_PORT") != "" {
		ConsulSetting.Port = os.Getenv("CONSUL_PORT")
	}
}

func mapTo(section string, v any) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		panic(err)
	}
}

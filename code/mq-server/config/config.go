package config

import (
	"fmt"

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
	mapTo("db", DatabaseSetting)
	mapTo("rabbitmq", RabbitMQSetting)

	DatabaseSetting.Url = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DatabaseSetting.User,
		DatabaseSetting.Password,
		DatabaseSetting.Host,
		DatabaseSetting.Port,
		DatabaseSetting.Name,
	)

	RabbitMQSetting.Url = fmt.Sprintf("%s://%s:%s@%s:%s/",
		RabbitMQSetting.RabbitMQ,
		RabbitMQSetting.RabbitMQUser,
		RabbitMQSetting.RabbitMQPassWord,
		RabbitMQSetting.RabbitMQHost,
		RabbitMQSetting.RabbitMQPort,
	)
}

func mapTo(section string, v any) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		panic(err)
	}
}

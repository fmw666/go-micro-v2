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
	mapTo("db", DatabaseSetting)
	mapTo("rabbitmq", RabbitMQSetting)

	// 从环境变量中读取 db 配置
	if os.Getenv("MYSQL_HOST") != "" {
		DatabaseSetting.Host = os.Getenv("MYSQL_HOST")
	}
	if os.Getenv("MYSQL_PORT") != "" {
		DatabaseSetting.Port = os.Getenv("MYSQL_PORT")
	}
	if os.Getenv("MYSQL_USERNAME") != "" {
		DatabaseSetting.User = os.Getenv("MYSQL_USERNAME")
	}
	if os.Getenv("MYSQL_PASSWORD") != "" {
		DatabaseSetting.Password = os.Getenv("MYSQL_PASSWORD")
	}
	if os.Getenv("MYSQL_DATABASE") != "" {
		DatabaseSetting.Name = os.Getenv("MYSQL_DATABASE")
	}

	// 从环境变量中读取 rabbitmq 配置
	if os.Getenv("RABBITMQ_HOST") != "" {
		RabbitMQSetting.RabbitMQHost = os.Getenv("RABBITMQ_HOST")
	}
	if os.Getenv("RABBITMQ_PORT") != "" {
		RabbitMQSetting.RabbitMQPort = os.Getenv("RABBITMQ_PORT")
	}
	if os.Getenv("RABBITMQ_USER") != "" {
		RabbitMQSetting.RabbitMQUser = os.Getenv("RABBITMQ_USER")
	}
	if os.Getenv("RABBITMQ_PASSWORD") != "" {
		RabbitMQSetting.RabbitMQPassWord = os.Getenv("RABBITMQ_PASSWORD")
	}

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

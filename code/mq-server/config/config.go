package config

import (
	"fmt"

	"github.com/go-ini/ini"
)

type Server struct {
	RunMode  string
	HttpPort string
}

var ServerSetting = &Server{}

type Database struct {
	Type     string
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Url      string
}

var DatabaseSetting = &Database{}

type RabbitMQ struct {
	RabbitMQ         string
	RabbitMQUser     string
	RabbitMQPassWord string
	RabbitMQHost     string
	RabbitMQPort     string
}

var RabbitMQSetting = &RabbitMQ{}

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
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		panic(err)
	}
}

package config

// db 配置
type Database struct {
	Type     string
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Url      string
}

// rabbitmq 配置
type RabbitMQ struct {
	RabbitMQ         string
	RabbitMQUser     string
	RabbitMQPassWord string
	RabbitMQHost     string
	RabbitMQPort     string
	Url              string
}

var DatabaseSetting = &Database{}
var RabbitMQSetting = &RabbitMQ{}

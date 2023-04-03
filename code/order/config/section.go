package config

// server 配置
type Server struct {
	Host             string
	RpcPort          string
	MicroServiceName string
}

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

// consul 配置
type Consul struct {
	Host string
	Port string
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

var ServerSetting = &Server{}
var DatabaseSetting = &Database{}
var ConsulSetting = &Consul{}
var RabbitMQSetting = &RabbitMQ{}

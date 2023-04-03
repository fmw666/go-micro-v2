package config

// server 配置
type Server struct {
	Host             string
	Port             string
	RpcPort          string
	MicroServiceName string
}

// consul 配置
type Consul struct {
	Host string
	Port string
}

// service 配置
type Service struct {
	OrderServiceName string
	UserServiceName  string
}

var ServerSetting = &Server{}
var ConsulSetting = &Consul{}
var ServiceSetting = &Service{}

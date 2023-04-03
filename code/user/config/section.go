package config

// server 配置
type Server struct {
	Host             string
	Port             string
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

var ServerSetting = &Server{}
var DatabaseSetting = &Database{}
var ConsulSetting = &Consul{}

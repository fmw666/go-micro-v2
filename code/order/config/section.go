package config

// server 配置
type Server struct {
	Host             string
	Port             string
	MicroServiceName string
}

// app 配置
type App struct {
	DefaultOffset string
	DefaultLimit  string
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

var ServerSetting = &Server{}
var AppSetting = &App{}
var DatabaseSetting = &Database{}

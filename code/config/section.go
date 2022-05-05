package config

// app 配置
type App struct {
	DefaultOffset string
	DefaultLimit  string
}

var AppSetting = &App{}

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

var DatabaseSetting = &Database{}

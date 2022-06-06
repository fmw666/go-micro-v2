package config

import "time"

// server 配置
type Server struct {
	Host string
	Port string
}

// app 配置
type App struct {
	DefaultOffset string
	DefaultLimit  string
	TimeFormat    string
}

// log 配置
type Log struct {
	SavePath   string
	SaveName   string
	FileExt    string
	TimeFormat string
}

// jwt 配置
type Jwt struct {
	Secret string
	Expire time.Duration
	Issuer string
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
var LogSetting = &Log{}
var JwtSetting = &Jwt{}
var DatabaseSetting = &Database{}

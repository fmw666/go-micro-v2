package main

import (
	"time"
	"user/config"
	"user/models"
	"user/pkg/logger"
	"user/pkg/utils/consul"
	"user/router"

	"github.com/micro/go-micro/v2/web"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @host localhost:8081
// @BasePath /api/v1
// @securityDefinitions.basic BasicAuth
// @in header
// @name Authorization
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	logger.Info("order service start...")

	// 初始化数据库
	models.Migrate()

	// gin Router 路由引擎
	ginRouter := router.Router()

	// consul 注册件
	consulReg := consul.ConsulReg

	// 获取一个微服务的实例
	microService := web.NewService(
		web.Name(config.ServerSetting.MicroServiceName),
		// 设置注册服务过期时间
		web.RegisterTTL(time.Second*30),
		//设置间隔多久再次注册服务
		web.RegisterInterval(time.Second*20),
		web.Address(config.ServerSetting.Host+":"+config.ServerSetting.Port),
		web.Handler(ginRouter),
		web.Registry(consulReg),
	)
	// 启动微服务
	if err := microService.Run(); err != nil {
		panic(err)
	}
}

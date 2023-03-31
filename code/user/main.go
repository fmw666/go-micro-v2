package main

import (
	"time"
	"user/config"
	"user/models"
	"user/pkg/logger"
	"user/router"

	"github.com/micro/go-micro/v2/web"
)

// @title User API
// @version 1.0
// @description User mciro service.
// @host localhost:8081
// @BasePath /api/v1
// @securityDefinitions.basic BasicAuth
// @in header
// @name Authorization
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	logger.Info("user service start...")

	// 初始化数据库
	models.Migrate()

	// gin Router 路由引擎
	ginRouter := router.Router()

	// 获取一个微服务的实例
	microService := web.NewService(
		web.Name(config.ServerSetting.MicroServiceName),
		// 设置注册服务过期时间
		web.RegisterTTL(time.Second*30),
		// 设置间隔多久再次注册服务
		web.RegisterInterval(time.Second*20),
		web.Address(config.ServerSetting.Host+":"+config.ServerSetting.Port),
		web.Handler(ginRouter),
	)
	// 启动微服务
	microService.Run()
}

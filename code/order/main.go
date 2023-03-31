package main

import (
	"order/config"
	"order/models"
	"order/pkg/logger"
	"order/router"

	"time"

	"github.com/micro/go-micro/v2/web"
)

// @title Order API
// @version 1.0
// @description Order micro service.
// @host localhost:8082
// @BasePath /api/v1
func main() {
	logger.Info("order service start...")

	// 初始化数据库
	models.Migrate()

	// gin Router 路由引擎
	ginRouter := router.Router()

	// 获取一个微服务的实例
	microService := web.NewService(
		web.Name(config.ServerSetting.MicroServiceName),
		// 设置注册服务的过期时间
		web.RegisterTTL(time.Second*30),
		// 设置间隔多久再次注册服务
		web.RegisterInterval(time.Second*20),
		web.Address(config.ServerSetting.Host+":"+config.ServerSetting.Port),
		web.Handler(ginRouter),
	)
	// 启动微服务
	microService.Run()
}

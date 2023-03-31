package router

import (
	v1 "order/api/v1"
	"order/middleware"

	_ "order/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router(services map[string]any) *gin.Engine {
	ginRouter := gin.Default()

	// Swagger 配置
	ginRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 中间件
	ginRouter.Use(middleware.Logger(), middleware.Cors(), middleware.ErrorMiddleware(), middleware.InitMiddleware(services))

	// 路由规则
	apiv1 := ginRouter.Group("/api/v1")
	apiUser := apiv1.Group("/orders")
	{
		apiUser.GET("", v1.GetOrderList)
		apiUser.POST("", v1.CreateOrder)
	}

	return ginRouter
}

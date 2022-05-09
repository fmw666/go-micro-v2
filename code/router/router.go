package router

import (
	"app/middleware"

	_ "app/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	ginRouter := gin.Default()

	// Swagger 配置
	ginRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 中间件
	ginRouter.Use(middleware.Cors(), middleware.ErrorMiddleware())

	// 路由规则
	apiv1 := ginRouter.Group("/api/v1")
	setupPingRouter(apiv1)
	apiUser := apiv1.Group("/user")
	setupUserRouter(apiUser)
	apiOrder := apiv1.Group("/orders")
	setupOrderRouter(apiOrder)

	return ginRouter
}

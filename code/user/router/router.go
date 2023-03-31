package router

import (
	v1 "user/api/v1"
	"user/middleware"

	_ "user/docs"

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
	apiUser := apiv1.Group("/user")
	{
		apiUser.POST("/register", v1.UserRegister)
		apiUser.POST("/login", v1.UserLogin)
	}
	{
		apiUser.Use(middleware.Authorization())
		apiUser.GET("/orders", v1.GetUserOrderList)
		apiUser.POST("/orders", v1.UserOrderCreate)
	}

	return ginRouter
}

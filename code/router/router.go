package router

import (
	"app/middleware"

	_ "app/docs"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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

	// 设置 sessionId 的密钥
	store := cookie.NewStore([]byte("something-very-secret"))
	ginRouter.Use(sessions.Sessions("SESSIONID", store))

	// 路由规则
	apiv1 := ginRouter.Group("/api/v1")
	setupPingRouter(apiv1)
	apiUser := apiv1.Group("/user")
	setupUserRouter(apiUser)
	// 需要登录的路由
	apiOrder := apiv1.Group("/orders")
	apiOrder.Use(middleware.Authorization())
	setupOrderRouter(apiOrder)

	return ginRouter
}

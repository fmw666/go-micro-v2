package router

import (
	"app/middleware"
	"app/service"

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
	{
		apiv1.GET("ping", func(context *gin.Context) {
			context.JSON(200, "success")
		})
		// 用户服务
		apiv1.POST("/user/register", service.UserRegister)
		apiv1.POST("/user/login", service.UserLogin)

		// 需要登录保护
		apiAuthed := apiv1.Group("/")
		apiAuthed.Use(middleware.Authorization())
		{
			apiOrder := apiAuthed.Group("/orders")
			{
				apiOrder.GET("", service.GetOrderList)
				apiOrder.POST("", service.CreateOrder)
			}
		}
	}
	return ginRouter
}

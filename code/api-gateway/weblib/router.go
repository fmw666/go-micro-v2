package weblib

import (
	_ "api-gateway/docs"
	"api-gateway/weblib/handlers"
	"api-gateway/weblib/middleware"

	"github.com/gin-gonic/gin"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// 路由 /api/v1
func setupPingRouter(router *gin.RouterGroup) {
	router.GET("ping", func(context *gin.Context) {
		context.JSON(200, "success")
	})
}

// 路由 /api/v1/user
func setupUserRouter(router *gin.RouterGroup) {
	router.POST("/register", handlers.UserRegister)
	router.POST("/login", handlers.UserLogin)
	apiOrder := router.Group("/orders")
	setupUserOrderRouter(apiOrder)
}

// 路由 /api/v1/user/orders
func setupUserOrderRouter(router *gin.RouterGroup) {
	router.Use(middleware.Authorization())
	router.GET("", handlers.GetUserOrderList)
	router.POST("", handlers.UserOrderCreate)
}

// 路由 /api/v1/orders
func setupOrderRouter(router *gin.RouterGroup) {
	router.GET("", handlers.GetOrderList)
	router.POST("", handlers.CreateOrder)
}

func NewRouter(services map[string]any) *gin.Engine {
	ginRouter := gin.Default()

	// Swagger 配置
	ginRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 中间件
	ginRouter.Use(middleware.Cors(), middleware.InitMiddleware(services), middleware.ErrorMiddleware())

	// 路由规则
	apiv1 := ginRouter.Group("/api/v1")
	setupPingRouter(apiv1)
	apiUser := apiv1.Group("/user")
	setupUserRouter(apiUser)
	apiOrder := apiv1.Group("/orders")
	setupOrderRouter(apiOrder)

	return ginRouter
}

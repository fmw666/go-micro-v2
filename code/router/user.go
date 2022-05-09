package router

import (
	apiv1 "app/api/v1"
	"app/middleware"

	"github.com/gin-gonic/gin"
)

// 路由 /api/v1/user
func setupUserRouter(router *gin.RouterGroup) {
	router.POST("/register", apiv1.UserRegister)
	router.POST("/login", apiv1.UserLogin)
	apiOrder := router.Group("/orders")
	setupUserOrderRouter(apiOrder)
}

// 路由 /api/v1/user/orders
func setupUserOrderRouter(router *gin.RouterGroup) {
	router.Use(middleware.Authorization())
	router.GET("", apiv1.GetUserOrderList)
	router.POST("", apiv1.UserOrderCreate)
}

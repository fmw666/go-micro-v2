package router

import (
	"app/service"

	"github.com/gin-gonic/gin"
)

// 路由 /api/v1/orders
func setupOrderRouter(router *gin.RouterGroup) {
	router.GET("", service.GetOrderList)
	router.POST("", service.CreateOrder)
}

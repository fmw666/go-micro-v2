package router

import (
	apiv1 "app/api/v1"

	"github.com/gin-gonic/gin"
)

// 路由 /api/v1/orders
func setupOrderRouter(router *gin.RouterGroup) {
	router.GET("", apiv1.GetOrderList)
	router.POST("", apiv1.CreateOrder)
}

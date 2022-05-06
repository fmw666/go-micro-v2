package router

import "github.com/gin-gonic/gin"

// 路由 /api/v1
func setupPingRouter(router *gin.RouterGroup) {
	router.GET("ping", func(context *gin.Context) {
		context.JSON(200, "success")
	})
}

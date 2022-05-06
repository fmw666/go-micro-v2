package router

import (
	apiv1 "app/api/v1"

	"github.com/gin-gonic/gin"
)

// 路由 /api/v1/user
func setupUserRouter(router *gin.RouterGroup) {
	// 用户服务
	router.POST("/register", apiv1.UserRegister)
	router.POST("/login", apiv1.UserLogin)
}

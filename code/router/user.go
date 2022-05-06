package router

import (
	"app/service"

	"github.com/gin-gonic/gin"
)

// 路由 /api/v1/user
func setupUserRouter(router *gin.RouterGroup) {
	// 用户服务
	router.POST("/register", service.UserRegister)
	router.POST("/login", service.UserLogin)
}

package utils

import (
	"app/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// 获取当前登录用户
func GetCurrentUser(c *gin.Context) *models.User {
	session := sessions.Default(c)
	return session.Get("currentUser").(*models.User)
}

// 设置当前登录用户
func SetCurrentUser(c *gin.Context, user *models.User) {
	session := sessions.Default(c)
	session.Set("currentUser", user)
	session.Save()
}

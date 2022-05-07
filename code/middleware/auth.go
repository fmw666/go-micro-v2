package middleware

import (
	"app/models"
	"app/pkg/utils"

	"strings"

	"github.com/gin-gonic/gin"
)

func Authorization() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		authHeaders := ginCtx.GetHeader("Authorization")
		authHeaderArr := strings.Split(authHeaders, ", ")
		// 多个 Authorization 头部，满足一个即可
		for _, authHeader := range authHeaderArr {
			context := strings.Fields(authHeader)
			if len(context) != 2 {
				continue
			}
			authType := context[0]
			authValue := context[1]
			// 设置变量监听是否认证成功
			success := false
			switch authType {
			case "Basic":
				basicAuth(authValue, &success)
			case "Bearer":
				jwtAuth(authValue, &success)
			default:
				continue
			}
			// 如果认证成功，则结束函数
			if success {
				ginCtx.Next()
				return
			}
		}
		// 如果认证失败，则返回鉴权失败
		ginCtx.JSON(500, gin.H{
			"code": 401,
			"msg":  "鉴权失败",
		})
		ginCtx.Abort()
	}
}

// Basic auth 用户认证
func basicAuth(auth string, success *bool) {
	username, password, ok := utils.ParseBasicAuth(auth)
	if !ok {
		return
	}
	// 校验用户名和密码
	var user models.User
	if err := models.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return
	}
	if !user.CheckPassword(password) {
		return
	}
	// 认证成功
	*success = true
}

// JWT token 用户认证
func jwtAuth(auth string, success *bool) {
	if auth == "" {
		return
	}
	// 校验 token
	_, err := utils.ParseToken(auth)
	if err != nil {
		return
	}
	// 认证成功
	*success = true
}

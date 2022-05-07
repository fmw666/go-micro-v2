package middleware

import (
	"app/models"
	"app/pkg/e"
	"app/pkg/utils"
	"net/http"

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
				basicAuth(authValue, ginCtx, &success)
			case "Bearer":
				jwtAuth(authValue, ginCtx, &success)
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
		var code e.ErrorCode = e.ERROR_AUTH_BASE
		ginCtx.JSON(http.StatusUnauthorized, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
		})
		ginCtx.Abort()
	}
}

// Basic auth 用户认证
func basicAuth(auth string, ginCtx *gin.Context, success *bool) {
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
	// 设置用户变量
	ginCtx.Set("user", user)
}

// JWT token 用户认证
func jwtAuth(auth string, ginCtx *gin.Context, success *bool) {
	if auth == "" {
		return
	}
	// 校验 token
	claims, err := utils.ParseToken(auth)
	if err != nil {
		return
	}
	// 获取 user_id
	userID := claims.Id
	// 校验 user_id
	var user models.User
	if err := models.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return
	}
	// 认证成功
	*success = true
	// 设置用户变量
	ginCtx.Keys["user"] = user
}

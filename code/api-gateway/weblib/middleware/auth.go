package middleware

import (
	"api-gateway/pkg/utils"
	"api-gateway/service"
	"context"
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
		ginCtx.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"message": "Unauthorized",
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
	// 从 gin.Key 中取出服务实例
	userService := ginCtx.Keys["userService"].(service.UserService)
	// 构建 req
	req := service.UserLoginRequest{
		Username: username,
		Password: password,
	}
	userResp, err := userService.UserLogin(context.Background(), &req)
	if err != nil {
		return
	}
	if userResp.Code != 0 {
		return
	}
	// 认证成功
	*success = true
	// 设置用户变量
	ginCtx.Set("user_id", userResp.Data.ID)
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
	// 认证成功
	*success = true
	// 设置用户变量
	ginCtx.Keys["user_id"] = uint32(userID)
}

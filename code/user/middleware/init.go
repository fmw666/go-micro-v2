package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 接受服务实例，并存到 gin.Key 中
func InitMiddleware(service []interface{}) gin.HandlerFunc {
	return func(context *gin.Context) {
		// 将实例存在 gin.Keys 中
		context.Keys = make(map[string]interface{})
		context.Keys["userService"] = service[0]
		// context.Keys["orderService"] = service[1]
		context.Next()
	}
}

// 错误处理中间件
func ErrorMiddleware() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				ginCtx.JSON(http.StatusNotFound, gin.H{
					"code": 404,
					"msg":  fmt.Sprintf("%s", r),
				})
				ginCtx.Abort()
			}
		}()
		ginCtx.Next()
	}
}

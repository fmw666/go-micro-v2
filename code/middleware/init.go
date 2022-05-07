package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// 错误处理中间件
func ErrorMiddleware() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				ginCtx.JSON(200, gin.H{
					"code": 404,
					"msg":  fmt.Sprintf("%s", r),
				})
				ginCtx.Abort()
			}
		}()
		ginCtx.Next()
	}
}

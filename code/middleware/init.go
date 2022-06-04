package middleware

import (
	"app/pkg/e"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 错误处理中间件
func ErrorMiddleware() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				ginCtx.JSON(http.StatusNotFound, gin.H{
					"code":    e.ERROR_EXCEPTION,
					"message": e.GetMsg(e.ERROR_EXCEPTION),
					"details": fmt.Sprintf("%s", r),
				})
				ginCtx.Abort()
			}
		}()
		ginCtx.Next()
	}
}

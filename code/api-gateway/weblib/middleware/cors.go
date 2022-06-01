package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Cors 跨域配置
func Cors() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		method := ginCtx.Request.Method               // 请求方法
		origin := ginCtx.Request.Header.Get("Origin") // 请求头部
		if origin != "" {
			ginCtx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			// 这是允许访问所有域
			ginCtx.Header("Access-Control-Allow-Origin", "*")
			// 服务器支持的所有跨域请求的方法，为了避免浏览次请求的多次'预检'请求
			ginCtx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			// header 的类型
			ginCtx.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			// 允许跨域设置                                                                                                      可以返回其他子段
			ginCtx.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar") // 跨域关键设置 让浏览器可以解析
			ginCtx.Header("Access-Control-Max-Age", "172800")                                                                                                                                                           // 缓存请求信息 单位为秒
			ginCtx.Header("Access-Control-Allow-Credentials", "false")                                                                                                                                                  //  跨域请求是否需要带cookie信息 默认设置为true
			ginCtx.Set("content-type", "application/json")                                                                                                                                                              // 设置返回格式是json
		}
		// 放行所有 OPTIONS 方法
		if method == "OPTIONS" {
			ginCtx.JSON(http.StatusOK, "Options Request!")
		}
		// 处理请求
		ginCtx.Next()
	}
}

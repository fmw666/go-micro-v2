package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// 日志配置
func Logger() gin.HandlerFunc {
	// 获取当前日期
	date := time.Now().Format("2006-01-02")

	// 创建日志文件所在的目录
	err := os.MkdirAll("logs", os.ModePerm)
	if err != nil {
		panic(fmt.Sprintf("Failed to create log directory: %v", err))
	}

	// 创建日志文件
	file, err := os.OpenFile(fmt.Sprintf("logs/gin_%s.log", date), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("Failed to create log file: %v", err))
	}

	return func(c *gin.Context) {
		// 记录请求开始时间
		start := time.Now()

		// 处理请求
		c.Next()

		// 记录请求结束时间
		end := time.Now()

		// 计算请求耗时
		latency := end.Sub(start)

		// 获取请求信息
		method := c.Request.Method
		path := c.Request.URL.Path
		statusCode := c.Writer.Status()

		// 记录日志
		message := fmt.Sprintf("| %3d | %13v | %15s | %s %s",
			statusCode,
			latency,
			c.ClientIP(),
			method,
			path,
		)
		fmt.Fprintln(file, message)
	}
}

package weblib

import (
	"api-gateway/docs"
	"api-gateway/weblib/handlers"
	"api-gateway/weblib/middleware"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initSwagger() {
	docs.SwaggerInfo.Title = "API Gateway"
	docs.SwaggerInfo.Description = "网关 api 总路由"
	docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = config.ServerSetting.PrefixUrl
	docs.SwaggerInfo.BasePath = "/api/v1"
	// docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func NewRouter(service ...interface{}) *gin.Engine {
	ginRouter := gin.Default()

	// Swagger 配置
	go initSwagger()
	ginRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	ginRouter.Use(middleware.Cors(), middleware.InitMiddleware(service), middleware.ErrorMiddleware())
	store := cookie.NewStore([]byte("something-very-secret"))
	ginRouter.Use(sessions.Sessions("mysession", store))
	v1 := ginRouter.Group("/api/v1")
	{
		v1.GET("ping", func(context *gin.Context) {
			context.JSON(200, "success")
		})
		// 用户服务
		v1.POST("/user/register", handlers.UserRegister)
		v1.POST("/user/login", handlers.UserLogin)

		// 需要登录保护
		// authed := v1.Group("/")
		// authed.Use(middleware.JWT())
		// {
		// authed.GET("tasks", handlers.GetTaskList)
		// authed.POST("task", handlers.CreateTask)
		// authed.GET("task/:id", handlers.GetTaskDetail) // task_id
		// authed.PUT("task/:id", handlers.UpdateTask)    // task_id
		// authed.DELETE("task/:id", handlers.DeleteTask) // task_id
		// }
	}
	return ginRouter
}

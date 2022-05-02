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
	// docs.SwaggerInfo.SecurityDefinitions.ApiKeyAuth = &docs.ApiKeyAuth{
	// 	In:   "header",
	// 	Name: "Authorization",
	// 	Type: "apiKey",
	// }
	// docs.SwaggerInfo.Host = config.ServerSetting.PrefixUrl
	docs.SwaggerInfo.BasePath = "/api/v1"
	// docs.SwaggerInfo.Schemes = []string{"http", "https"}
}

func NewRouter(service ...interface{}) *gin.Engine {
	ginRouter := gin.Default()

	// Swagger 配置
	initSwagger()
	ginRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	ginRouter.Use(middleware.Cors(), middleware.InitMiddleware(service), middleware.ErrorMiddleware())
	store := cookie.NewStore([]byte("something-very-secret"))
	ginRouter.Use(sessions.Sessions("mysession", store))
	apiv1 := ginRouter.Group("/api/v1")
	{
		apiv1.GET("ping", func(context *gin.Context) {
			context.JSON(200, "success")
		})
		// 用户服务
		apiv1.POST("/user/register", handlers.UserRegister)
		apiv1.POST("/user/login", handlers.UserLogin)

		// 需要登录保护
		apiAuthed := apiv1.Group("/")
		apiAuthed.Use(middleware.JWT())
		{
			apiOrder := apiAuthed.Group("/orders")
			{
				apiOrder.GET("", handlers.GetOrderList)
				apiOrder.GET(":id", handlers.GetOrderDetail)
			}
		}
	}
	return ginRouter
}

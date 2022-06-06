package main

import (
	"app/config"
	"app/models"
	"app/router"
)

// @title 单体应用 swagger-ui API 文档
// @version 1.0
// @description 详情见源码地址：https://github.com/fmw666/microservice-code-sample/tree/monolithic-app.
// @host localhost:8080
// @BasePath /api/v1
// @securityDefinitions.basic BasicAuth
// @in header
// @name Authorization
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	models.Migrate()
	ginRouter := router.Router()
	ginRouter.Run(config.ServerSetting.Host + ":" + config.ServerSetting.Port)
}

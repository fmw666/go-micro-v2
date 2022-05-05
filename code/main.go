package main

import (
	"app/models"
	"app/router"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
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
	ginRouter.Run(":8080")
}

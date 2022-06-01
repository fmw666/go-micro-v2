package handlers

import (
	"api-gateway/service"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册
// @Summary 用户注册
// @Description 微服务模块 User 中提供的用户注册服务
// @Tags User 服务
// @Accept  json
// @Produce  json
// @Param body body schema.Register true "注册"
// @Success 200 {string} json "{"code":200,"data":{},"message":""}"
// @Router /user/register [post]
func UserRegister(ginCtx *gin.Context) {
	var userReq service.UserRegisterRequest
	PanicIfUserError(ginCtx.Bind(&userReq))
	// 从 gin.Key 中取出服务实例
	userService := ginCtx.Keys["userService"].(service.UserService)
	userResp, err := userService.UserRegister(context.Background(), &userReq)
	PanicIfUserError(err)

	ginCtx.JSON(http.StatusOK, gin.H{
		"code":    userResp.Code,
		"data":    userResp.Data,
		"message": userResp.Message,
	})
}

// UserLogin 用户登录
// @Summary 用户登录
// @Description 微服务模块 User 中提供的用户登录服务
// @Tags User 服务
// @Accept  json
// @Produce  json
// @Param body body schema.Login true "登录"
// @Success 200 {string} json "{"code":200,"data":{},"msg":""}"
// @Router /user/login [post]
func UserLogin(ginCtx *gin.Context) {
	var userReq service.UserLoginRequest
	PanicIfUserError(ginCtx.Bind(&userReq))
	// 从 gin.Key 中取出服务实例
	userService := ginCtx.Keys["userService"].(service.UserService)
	userResp, err := userService.UserLogin(context.Background(), &userReq)
	PanicIfUserError(err)

	ginCtx.JSON(http.StatusOK, gin.H{
		"code":    userResp.Code,
		"data":    userResp.Data,
		"message": userResp.Message,
	})
}

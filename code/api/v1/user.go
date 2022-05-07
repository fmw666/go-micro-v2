package v1

import (
	"app/pkg/utils"
	"app/service"

	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册
// @Summary 用户注册
// @Description User 服务中提供的用户注册服务
// @Tags User 服务
// @Accept  json
// @Produce  json
// @Param body body schema.RegisterReq true "注册"
// @Success 200 {string} json "{"code":0,"data":{}}"
// @Router /user/register [post]
func UserRegister(ginCtx *gin.Context) {
	data, code := service.UserRegister(ginCtx)
	utils.Response(ginCtx, code, data)
}

// UserLogin 用户登录
// @Summary 用户登录
// @Description User 服务中提供的用户登录服务
// @Tags User 服务
// @Accept  json
// @Produce  json
// @Param body body schema.LoginReq true "登录"
// @Success 200 {string} json "{"code":0,"data":{}}"
// @Router /user/login [post]
func UserLogin(ginCtx *gin.Context) {
	data, code := service.UserLogin(ginCtx)
	utils.Response(ginCtx, code, data)
}

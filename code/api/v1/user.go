package v1

import (
	"app/pkg/e"
	"app/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册
// @Summary 用户注册
// @Description User 服务中提供的用户注册服务
// @Tags User 服务
// @Accept  json
// @Produce  json
// @Param body body schema.RegisterReq true "注册"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /user/register [post]
func UserRegister(ginCtx *gin.Context) {
	data, code := service.UserRegister(ginCtx)
	if code != e.SUCCESS {
		ginCtx.JSON(http.StatusOK, gin.H{"code": code, "msg": e.GetMsg(code)})
		return
	}
	ginCtx.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "data": data})
}

// UserLogin 用户登录
// @Summary 用户登录
// @Description User 服务中提供的用户登录服务
// @Tags User 服务
// @Accept  json
// @Produce  json
// @Param body body schema.LoginReq true "登录"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /user/login [post]
func UserLogin(ginCtx *gin.Context) {
	data, code := service.UserLogin(ginCtx)
	if code != e.SUCCESS {
		ginCtx.JSON(http.StatusOK, gin.H{"code": code, "msg": e.GetMsg(code)})
		return
	}
	ginCtx.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "data": data})
}

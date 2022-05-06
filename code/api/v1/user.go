package v1

import (
	"app/schema"
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
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /user/register [post]
func UserRegister(ginCtx *gin.Context) {
	// 获取 body 内容
	var req schema.RegisterReq
	if err := ginCtx.ShouldBindJSON(&req); err != nil {
		ginCtx.JSON(200, gin.H{"code": 400, "msg": "请求参数错误"})
		return
	}
	data, err := service.UserRegister(req)
	if err != nil {
		ginCtx.JSON(200, gin.H{"code": 500, "msg": err.Error()})
		return
	}
	ginCtx.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": data,
	})
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
	// 获取 body 内容
	var loginReq schema.LoginReq
	if err := ginCtx.ShouldBindJSON(&loginReq); err != nil {
		ginCtx.JSON(200, gin.H{"code": 400, "msg": "请求参数错误"})
		return
	}
	data, err := service.UserLogin(loginReq)
	if err != nil {
		ginCtx.JSON(200, gin.H{"code": 500, "msg": err.Error()})
		return
	}
	ginCtx.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": data,
	})
}

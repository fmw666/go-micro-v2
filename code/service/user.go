package service

import (
	"app/models"
	"app/pkg/utils"
	"app/schema"

	"github.com/gin-gonic/gin"
)

func BuildUser(user models.User) *schema.UserResp {
	return &schema.UserResp{
		ID:        uint(user.Id),
		Username:  user.Username,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

// UserRegister 用户注册
// @Summary 用户注册
// @Description 微服务模块 User 中提供的用户注册服务
// @Tags User 服务
// @Accept  json
// @Produce  json
// @Param body body schema.RegisterReq true "注册"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /user/register [post]
func UserRegister(ginCtx *gin.Context) {
	// 获取 body 内容
	var registerReq schema.RegisterReq
	if err := ginCtx.ShouldBindJSON(&registerReq); err != nil {
		ginCtx.JSON(200, gin.H{"code": 400, "msg": "请求参数错误"})
		return
	}
	// 校验用户名是否已经存在
	var count int64 = 0
	if err := models.DB.Model(&models.User{}).Where("username = ?", registerReq.Username).Count(&count).Error; err != nil {
		ginCtx.JSON(200, gin.H{"code": 400, "msg": "数据库查询错误"})
		return
	}
	if count > 0 {
		ginCtx.JSON(200, gin.H{"code": 400, "msg": "用户名已存在"})
		return
	}
	// 创建用户
	user := models.User{
		Username: registerReq.Username,
	}
	// 加密密码
	if err := user.SetPassword(registerReq.Password); err != nil {
		ginCtx.JSON(200, gin.H{"code": 400, "msg": "密码加密错误"})
		return
	}
	// 插入数据库
	if err := models.DB.Create(&user).Error; err != nil {
		ginCtx.JSON(200, gin.H{"code": 400, "msg": "数据库创建错误"})
		return
	}
	ginCtx.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": gin.H{
			"user": BuildUser(user),
		},
	})
}

// UserLogin 用户登录
// @Summary 用户登录
// @Description 微服务模块 User 中提供的用户登录服务
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
	// 判断用户名是否存在
	var user models.User
	if err := models.DB.Where("username = ?", loginReq.Username).First(&user).Error; err != nil {
		ginCtx.JSON(200, gin.H{"code": 400, "msg": "用户名不存在"})
		return
	}
	// 校验密码
	if !user.CheckPassword(loginReq.Password) {
		ginCtx.JSON(200, gin.H{"code": 400, "msg": "密码错误"})
		return
	}
	// 获取 token
	token, _ := utils.GenerateToken(uint(user.Id))
	ginCtx.JSON(200, gin.H{
		"code": 200,
		"msg":  "ok",
		"data": gin.H{
			"user":  BuildUser(user),
			"token": token,
		},
	})
}

package service

import (
	"app/models"
	"app/pkg/e"
	"app/pkg/utils"
	"app/schema"

	"github.com/gin-gonic/gin"
)

func buildUser(user models.User) *schema.UserResp {
	return &schema.UserResp{
		ID:        uint(user.Id),
		Username:  user.Username,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func UserRegister(ginCtx *gin.Context) (*schema.UserResp, e.ErrorCode) {
	// 获取 body 内容
	var req schema.RegisterReq
	if err := ginCtx.ShouldBindJSON(&req); err != nil {
		return nil, e.ERROR_PARAM_INVALID
	}
	// 判断密码是否一致
	if req.Password != req.PasswordConfirm {
		return nil, e.ERROR_PASSWORD_NOT_MATCH
	}
	// 校验用户名是否已经存在
	var count int64 = 0
	if err := models.DB.Model(&models.User{}).Where("username = ?", req.Username).Count(&count).Error; err != nil {
		// 数据库查询错误
		return nil, e.ERROR_DB_BASE
	}
	// 判断用户名是否已经存在
	if count > 0 {
		return nil, e.ERROR_USER_EXIST
	}
	// 创建用户
	user := models.User{
		Username: req.Username,
	}
	// 加密密码
	if err := user.SetPassword(req.Password); err != nil {
		return nil, e.ERROR_USER_SET_PASSWORD
	}
	// 插入数据库
	if err := models.DB.Create(&user).Error; err != nil {
		return nil, e.ERROR_DB_CREATE
	}
	// 返回用户信息
	return buildUser(user), e.SUCCESS
}

func UserLogin(ginCtx *gin.Context) (interface{}, e.ErrorCode) {
	// 获取 body 内容
	var req schema.LoginReq
	if err := ginCtx.ShouldBindJSON(&req); err != nil {
		return nil, e.ERROR_PARAM_INVALID
	}
	// 判断用户名是否存在
	var user models.User
	if err := models.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		return nil, e.ERROR_USER_NOT_FOUND
	}
	// 校验密码
	if !user.CheckPassword(req.Password) {
		return nil, e.ERROR_USER_PASSWORD
	}
	// 获取 token
	token, _ := utils.GenerateToken(uint(user.Id))

	data := gin.H{
		"token": token,
		"user":  buildUser(user),
	}
	return data, e.SUCCESS
}

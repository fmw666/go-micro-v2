package service

import (
	"user/models"
	"user/pkg/e"
	"user/pkg/utils"
	"user/schema"

	"github.com/gin-gonic/gin"
)

func buildUser(user models.User) *schema.UserResp {
	return &schema.UserResp{
		ID:        user.Id,
		Username:  user.Username,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func UserRegister(username, password, passwordConfirm string) (*schema.UserResp, e.ErrorCode) {
	// 判断密码是否一致
	if password != passwordConfirm {
		return nil, e.ERROR_PASSWORD_NOT_MATCH
	}
	// 校验用户名是否已经存在
	var count int64 = 0
	if err := models.DB.Model(&models.User{}).Where("username = ?", username).Count(&count).Error; err != nil {
		// 数据库查询错误
		return nil, e.ERROR_DB_BASE
	}
	// 判断用户名是否已经存在
	if count > 0 {
		return nil, e.ERROR_USER_EXIST
	}
	// 创建用户
	user := models.User{
		Username: username,
	}
	// 加密密码
	if err := user.SetPassword(password); err != nil {
		return nil, e.ERROR_USER_SET_PASSWORD
	}
	// 插入数据库
	if err := models.DB.Create(&user).Error; err != nil {
		return nil, e.ERROR_DB_CREATE
	}
	// 返回用户信息
	return buildUser(user), e.SUCCESS
}

func UserLogin(username, password string) (interface{}, e.ErrorCode) {
	// 判断用户名是否存在
	var user models.User
	if err := models.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, e.ERROR_USER_NOT_FOUND
	}
	// 校验密码
	if !user.CheckPassword(password) {
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

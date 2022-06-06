package service

import (
	"app/config"
	"app/models"
	"app/pkg/e"
	"app/pkg/utils"
	"app/schema"
	userSchema "app/schema/user"
)

func buildUser(user models.User, token string) *userSchema.UserDetailWithToken {
	return &userSchema.UserDetailWithToken{
		Token: token,
		User: userSchema.UserDetail{
			ID:        user.Id,
			Username:  user.Username,
			CreatedAt: user.CreatedAt.Format(config.AppSetting.TimeFormat),
			UpdatedAt: user.UpdatedAt.Format(config.AppSetting.TimeFormat),
		},
	}
}

func UserRegister(username, password, passwordConfirm string) (resp schema.Response) {
	// 判断密码是否一致
	if password != passwordConfirm {
		resp.Code = e.ERROR_PASSWORD_NOT_MATCH
		resp.Message = e.GetMsg(e.ERROR_PASSWORD_NOT_MATCH)
		return
	}
	// 校验用户名是否已经存在
	var count int64 = 0
	if err := models.DB.Model(&models.User{}).Where("username = ?", username).Count(&count).Error; err != nil {
		// 数据库查询错误
		resp.Code = e.ERROR_DB_BASE
		resp.Message = e.GetMsg(e.ERROR_DB_BASE)
		return
	}
	// 判断用户名是否已经存在
	if count > 0 {
		resp.Code = e.ERROR_USER_EXIST
		resp.Message = e.GetMsg(e.ERROR_USER_EXIST)
		return
	}
	// 创建用户
	user := models.User{
		Username: username,
	}
	// 加密密码
	if err := user.SetPassword(password); err != nil {
		resp.Code = e.ERROR_USER_SET_PASSWORD
		resp.Message = e.GetMsg(e.ERROR_USER_SET_PASSWORD)
		return
	}
	// 插入数据库
	if err := models.DB.Create(&user).Error; err != nil {
		resp.Code = e.ERROR_DB_BASE
		resp.Message = e.GetMsg(e.ERROR_DB_BASE)
		return
	}
	// 获取 token
	token, _ := utils.GenerateToken(uint(user.Id))

	// 返回用户信息
	resp.Code = e.SUCCESS
	resp.Data = buildUser(user, token)
	return
}

func UserLogin(username, password string) (resp schema.Response) {
	// 判断用户名是否存在
	var user models.User
	if err := models.DB.Where("username = ?", username).First(&user).Error; err != nil {
		resp.Code = e.ERROR_USER_NOT_FOUND
		resp.Message = e.GetMsg(e.ERROR_USER_NOT_FOUND)
		return
	}
	// 校验密码
	if !user.CheckPassword(password) {
		resp.Code = e.ERROR_USER_PASSWORD
		resp.Message = e.GetMsg(e.ERROR_USER_PASSWORD)
		return
	}
	// 获取 token
	token, _ := utils.GenerateToken(uint(user.Id))

	resp.Code = e.SUCCESS
	resp.Data = buildUser(user, token)
	return
}

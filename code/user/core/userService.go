package core

import (
	"context"
	"user/models"
	"user/pkg/e"
	"user/pkg/utils"
	"user/service"
)

func (*UserService) UserLogin(ctx context.Context, req *service.UserLoginRequest, resp *service.UserDetailResponse) error {
	// 参数校验
	if req.Username == "" || req.Password == "" {
		resp.Code = e.ERROR_PARAM_NOT_CONTENT
		resp.Message = e.GetMsg(e.ERROR_PARAM_NOT_CONTENT)
		return nil
	}
	var user models.User
	// 查询用户是否存在
	if err := models.DB.Where("username=?", req.Username).First(&user).Error; err != nil {
		resp.Code = e.ERROR_USER_NOT_FOUND
		resp.Message = e.GetMsg(e.ERROR_USER_NOT_FOUND)
		return nil
	}
	// 校验密码
	if !user.CheckPassword(req.Password) {
		resp.Code = e.ERROR_USER_PASSWORD
		resp.Message = e.GetMsg(e.ERROR_USER_PASSWORD)
		return nil
	}
	// 获取 token
	token, _ := utils.GenerateToken(user.Id)
	// 生成响应
	resp.Code = e.SUCCESS
	resp.Data = &service.UserResponse{
		User:  buildUser(user),
		Token: token,
	}
	return nil
}

func (*UserService) UserRegister(ctx context.Context, req *service.UserRegisterRequest, resp *service.UserDetailResponse) error {
	// 参数校验
	if req.Username == "" || req.Password == "" || req.PasswordConfirm == "" {
		resp.Code = e.ERROR_PARAM_NOT_CONTENT
		resp.Message = e.GetMsg(e.ERROR_PARAM_NOT_CONTENT)
		return nil
	}
	if req.Password != req.PasswordConfirm {
		resp.Code = e.ERROR_PASSWORD_NOT_MATCH
		resp.Message = e.GetMsg(e.ERROR_PASSWORD_NOT_MATCH)
		return nil
	}
	// 查询是否存在同名用户
	var count int64 = 0
	if err := models.DB.Model(&models.User{}).Where("username=?", req.Username).Count(&count).Error; err != nil {
		resp.Code = e.ERROR_DB_BASE
		resp.Message = e.GetMsg(e.ERROR_DB_BASE)
		return nil
	}
	if count > 0 {
		resp.Code = e.ERROR_USER_EXIST
		resp.Message = e.GetMsg(e.ERROR_USER_EXIST)
		return nil
	}
	user := models.User{Username: req.Username}
	// 加密密码
	if err := user.SetPassword(req.Password); err != nil {
		resp.Code = e.ERROR_USER_SET_PASSWORD
		resp.Message = e.GetMsg(e.ERROR_USER_SET_PASSWORD)
		return nil
	}
	// 创建用户
	if err := models.DB.Create(&user).Error; err != nil {
		resp.Code = e.ERROR_DB_BASE
		resp.Message = e.GetMsg(e.ERROR_DB_BASE)
		return nil
	}
	// 获取 token
	token, _ := utils.GenerateToken(user.Id)
	// 生成响应
	resp.Code = e.SUCCESS
	resp.Data = &service.UserResponse{
		User:  buildUser(user),
		Token: token,
	}
	return nil
}

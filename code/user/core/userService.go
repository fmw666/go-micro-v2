package core

import (
	"context"
	"fmt"
	"user/models"
	"user/pkg/e"
	"user/schema"
	"user/service"
)

func (*UserService) UserLogin(ctx context.Context, req *service.UserLoginRequest, resp *service.UserDetailResponse) error {
	if req.Username == "" || req.Password == "" {
		resp.Code = e.ERROR_PARAM_NOT_CONTENT
		return nil
	}
	var user models.User
	if err := models.DB.Where("username=?", req.Username).First(&user).Error; err != nil {
		resp.Code = e.ERROR_USER_NOT_FOUND
		return nil
	}
	if !user.CheckPassword(req.Password) {
		resp.Code = e.ERROR_USER_PASSWORD
		return nil
	}
	resp.Code = e.SUCCESS
	resp.UserDetail = schema.EncodeUser(user)
	return nil
}

func (*UserService) UserRegister(ctx context.Context, req *service.UserRegisterRequest, resp *service.UserDetailResponse) error {
	// TODO 为什么是 passwordconfirm 而非 password_confirm
	fmt.Println("UserRegister")
	fmt.Println(req.Password)
	fmt.Println(req.PasswordConfirm)
	if req.Password != req.PasswordConfirm {
		resp.Code = e.ERROR_PASSWORD_NOT_MATCH
		return nil
	}
	var count int64 = 0
	if err := models.DB.Model(&models.User{}).Where("username=?", req.Username).Count(&count).Error; err != nil {
		resp.Code = e.ERROR_DB_BASE
		return nil
	}
	if count > 0 {
		resp.Code = e.ERROR_USER_EXIST
		return nil
	}
	user := models.User{Username: req.Username}
	// 加密密码
	if err := user.SetPassword(req.Password); err != nil {
		resp.Code = e.ERROR_USER_SET_PASSWORD
		return nil
	}
	if err := models.DB.Create(&user).Error; err != nil {
		resp.Code = e.ERROR_DB_CREATE
		return nil
	}
	resp.Code = e.SUCCESS
	resp.UserDetail = schema.EncodeUser(user)
	return nil
}

package core

import (
	"context"
	"errors"
	"net/http"
	"user/models"
	"user/schema"
	"user/service"
)

func (*UserService) UserLogin(ctx context.Context, req *service.UserLoginRequest, resp *service.UserDetailResponse) error {
	if req.Username == "" || req.Password == "" {
		// return errors.New("用户名或密码不能为空")
		resp.Code = http.StatusNoContent
		return nil
	}
	var user models.User
	resp.Code = 200
	if err := models.DB.Where("username=?", req.Username).First(&user).Error; err != nil {
		resp.Code = 400
		return nil
	}
	if !user.CheckPassword(req.Password) {
		resp.Code = 400
		return nil
	}
	resp.UserDetail = schema.EncodeUser(user)
	return nil
}

func (*UserService) UserRegister(ctx context.Context, req *service.UserRegisterRequest, resp *service.UserDetailResponse) error {
	if req.Password != req.PasswordConfirm {
		err := errors.New("两次密码输入不一致")
		return err
	}
	var count int64 = 0
	if err := models.DB.Model(&models.User{}).Where("username=?", req.Username).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		err := errors.New("用户名已存在")
		return err
	}
	user := models.User{
		Username: req.Username,
	}
	// 加密密码
	if err := user.SetPassword(req.Password); err != nil {
		return err
	}
	if err := models.DB.Create(&user).Error; err != nil {
		return err
	}
	resp.UserDetail = schema.EncodeUser(user)
	return nil
}

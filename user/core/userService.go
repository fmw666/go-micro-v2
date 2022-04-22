package core

import (
	"context"
	"errors"
	"user/models"
	"user/service"

	"gorm.io/gorm"
)

func BuildUser(item models.User) *service.UserModel {
	userModel := service.UserModel{
		// ID:        uint32(item.ID),
		Username:  item.Username,
		CreatedAt: item.CreatedAt.Unix(),
		UpdatedAt: item.UpdatedAt.Unix(),
	}
	return &userModel
}

func (*UserService) UserLogin(ctx context.Context, req *service.UserRequest, resp *service.UserDetailResponse) error {
	var user models.User
	resp.Code = 200
	if err := models.DB.Where("user_name=?", req.Username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			resp.Code = 400
			return nil
		}
		resp.Code = 500
		return nil
	}
	if !user.CheckPassword(req.Password) {
		resp.Code = 400
		return nil
	}
	resp.UserDetail = BuildUser(user)
	return nil
}

func (*UserService) UserRegister(ctx context.Context, req *service.UserRequest, resp *service.UserDetailResponse) error {
	if req.Password != req.PasswordConfirm {
		err := errors.New("两次密码输入不一致")
		return err
	}
	var count int64 = 0
	if err := models.DB.Model(&models.User{}).Where("user_name=?", req.Username).Count(&count).Error; err != nil {
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
	resp.UserDetail = BuildUser(user)
	return nil
}

package core

import (
	"context"
	"errors"
	"fmt"
	"user/models"
	"user/service"

	"github.com/jinzhu/gorm"
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
	fmt.Println("User login...")
	if err := models.DB.Where("username=?", req.Username).First(&user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			fmt.Println("User not found...")
			resp.Code = 400
			return nil
		}
		fmt.Println("User login error...")
		resp.Code = 500
		return nil
	}
	if !user.CheckPassword(req.Password) {
		fmt.Println("User password error...")
		resp.Code = 400
		return nil
	}
	resp.UserDetail = BuildUser(user)
	fmt.Println("User login success...")
	return nil
}

func (*UserService) UserRegister(ctx context.Context, req *service.UserRequest, resp *service.UserDetailResponse) error {
	fmt.Println("User register...")
	fmt.Println(req)
	// if req.Password != req.PasswordConfirm {
	// 	err := errors.New("两次密码输入不一致")
	// 	return err
	// }
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
	resp.UserDetail = BuildUser(user)
	return nil
}

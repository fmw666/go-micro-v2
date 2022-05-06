package service

import (
	"app/models"
	"app/pkg/utils"
	"app/schema"
	"errors"

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

func UserRegister(req schema.RegisterReq) (*schema.UserResp, error) {
	// 校验用户名是否已经存在
	var count int64 = 0
	if err := models.DB.Model(&models.User{}).Where("username = ?", req.Username).Count(&count).Error; err != nil {
		// 数据库查询错误
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("用户名已经存在")
	}
	// 创建用户
	user := models.User{
		Username: req.Username,
	}
	// 加密密码
	if err := user.SetPassword(req.Password); err != nil {
		return nil, err
	}
	// 插入数据库
	if err := models.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	// 返回用户信息
	return buildUser(user), nil
}

func UserLogin(req schema.LoginReq) (interface{}, error) {
	// 判断用户名是否存在
	var user models.User
	if err := models.DB.Where("username = ?", req.Username).First(&user).Error; err != nil {
		return nil, errors.New("用户名不存在")
	}
	// 校验密码
	if !user.CheckPassword(req.Password) {
		return nil, errors.New("密码错误")
	}
	// 获取 token
	token, _ := utils.GenerateToken(uint(user.Id))

	data := gin.H{
		"token": token,
		"user":  buildUser(user),
	}
	return data, nil
}

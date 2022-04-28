package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID
	Timestamps
	SoftDelete
	Username string `gorm:"column:username;type:varchar(255);unique;" json:"username"`
	Password string `gorm:"column:password;type:varchar(255);" json:"password"`
}

const (
	PassWordCost = 12 // 密码加密难度
)

// 表名
func (user *User) TableName() string {
	return "user"
}

// 加密密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// 检验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

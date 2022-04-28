package models

import (
	"log"

	"order/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB = Init()

func Init() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.DatabaseSetting.Url), &gorm.Config{})
	if err != nil {
		log.Println("gorm Init Error : ", err)
	}
	return db
}

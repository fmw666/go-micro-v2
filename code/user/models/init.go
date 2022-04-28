package models

import (
	"log"

	"user/config"

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

func Migrate() {
	err := DB.Set(`gorm:table_options`, "charset=utf8mb4").AutoMigrate(&User{})
	if err != nil {
		log.Println("gorm Init Error : ", err)
	}
}

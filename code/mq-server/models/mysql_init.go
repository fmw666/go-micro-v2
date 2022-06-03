package models

import (
	"mq-server/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// DB 数据库连接单例
var DB = Init_DB(config.DatabaseSetting.Url)

func Init_DB(connString string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
		// 默认表名不加复数
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()

	// 设置连接池
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Second * 30)

	return db
}

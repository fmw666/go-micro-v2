package models

import (
	"time"

	"gorm.io/gorm"
)

// 自增 ID 主键
type ID struct {
	Id uint32 `json:"id" gorm:"primaryKey"`
}

// 创建、更新时间
type Timestamps struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// 软删除
type SoftDelete struct {
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

package models

type Order struct {
	ID
	Timestamps
	SoftDelete
	Name   string `gorm:"column:name;type:varchar(255);unique;" json:"name"`
	UserID uint   `gorm:"column:user_id;type:int(11);" json:"user_id"`
}

// 表名
func (order *Order) TableName() string {
	return "order"
}

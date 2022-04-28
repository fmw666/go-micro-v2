package models

type Order struct {
	ID
	Timestamps
	SoftDelete
	Name string `gorm:"column:name;type:varchar(255);unique;" json:"name"`
}

// 表名
func (order *Order) TableName() string {
	return "order"
}

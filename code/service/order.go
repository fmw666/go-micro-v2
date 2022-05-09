package service

import (
	"app/models"
	"app/pkg/e"
	"app/schema"
)

// 订单详情
func buildOrder(order models.Order) *schema.OrderResp {
	return &schema.OrderResp{
		ID:        order.Id,
		Name:      order.Name,
		UserID:    order.UserID,
		CreatedAt: order.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: order.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

// 订单列表
func buildOrderList(orders []*models.Order) []*schema.OrderResp {
	data := make([]*schema.OrderResp, 0)
	for _, order := range orders {
		data = append(data, buildOrder(*order))
	}
	return data
}

// 获取订单列表
func GetOrderList(offset, limit int, userID ...uint32) (int64, []*schema.OrderResp, e.ErrorCode) {
	var count int64
	data := make([]*models.Order, 0)
	// 查询订单列表
	orders := models.DB.Model(new(models.Order))
	if len(userID) > 0 {
		orders = orders.Where("user_id = ?", userID[0])
	}
	err := orders.Count(&count).Offset(offset).Limit(limit).Order("id desc").Find(&data).Error
	if err != nil {
		return 0, nil, e.ERROR_DB_BASE
	}
	return count, buildOrderList(data), e.SUCCESS
}

// 创建订单
func CreateOrder(name string, userID uint32) (*schema.OrderResp, e.ErrorCode) {
	// 创建订单模型
	order := &models.Order{
		Name:   name,
		UserID: userID,
	}

	err := models.DB.Create(order).Error
	if err != nil {
		return nil, e.ERROR_DB_BASE
	}

	return buildOrder(*order), e.SUCCESS
}

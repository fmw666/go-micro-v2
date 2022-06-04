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
func GetOrderList(offset, limit, userID uint32) (resp schema.Response) {
	var count int64
	data := make([]*models.Order, 0)
	// 查询订单列表
	orders := models.DB.Model(new(models.Order))
	if userID != 0 {
		orders = orders.Where("user_id = ?", userID)
	}
	err := orders.Count(&count).Offset(int(offset)).Limit(int(limit)).Order("id desc").Find(&data).Error
	if err != nil {
		resp.Code = e.ERROR_DB_BASE
		resp.Message = e.GetMsg(e.ERROR_DB_BASE)
		return
	}
	// 生成响应
	resp.Code = e.SUCCESS
	resp.Data = buildOrderList(data)
	resp.PageInfo = &schema.PageInfo{
		Offset: offset,
		Limit:  limit,
		Total:  uint32(count),
	}
	return
}

// 创建订单
func CreateOrder(name string, userID uint32) (resp schema.Response) {
	// 创建订单模型
	order := &models.Order{
		Name:   name,
		UserID: userID,
	}

	err := models.DB.Create(order).Error
	if err != nil {
		resp.Code = e.ERROR_DB_BASE
		resp.Message = e.GetMsg(e.ERROR_DB_BASE)
		return
	}

	resp.Code = e.SUCCESS
	resp.Data = buildOrder(*order)
	return
}

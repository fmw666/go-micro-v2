package service

import (
	"app/models"
	"app/schema"
)

func buildOrder(order models.Order) *schema.OrderResp {
	return &schema.OrderResp{
		ID:        uint(order.Id),
		Name:      order.Name,
		UserID:    uint(order.UserID),
		CreatedAt: order.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: order.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

func buildOrderList(orders []*models.Order) []*schema.OrderResp {
	data := make([]*schema.OrderResp, 0)
	for _, order := range orders {
		data = append(data, buildOrder(*order))
	}
	return data
}

func GetOrderList(offset, limit int, userID int) (int64, []*schema.OrderResp, error) {
	var count int64
	data := make([]*models.Order, 0)
	// 查询订单列表
	orders := models.DB.Model(new(models.Order))
	if userID != 0 {
		orders = orders.Where("user_id = ?", userID)
	}
	err := orders.Count(&count).Offset(offset).Limit(limit).Order("id desc").Find(&data).Error
	return count, buildOrderList(data), err
}

func CreateOrder(req schema.OrderCreateReq) (*schema.OrderResp, error) {
	// 获取当前登录用户
	// user := utils.GetCurrentUser(ginCtx)
	// logger.Info("user: %+v", user)

	order := &models.Order{
		Name: req.Name,
		// UserID: uint(user.Id),
	}

	err := models.DB.Create(order).Error
	if err != nil {
		return nil, err
	}

	return buildOrder(*order), nil
}

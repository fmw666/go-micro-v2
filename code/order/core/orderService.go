package core

import (
	"context"
	"order/models"
	"order/pkg/e"
	"order/schema"
	"order/service"
)

// 微服务调用的 GetOrderList 方法
func (*OrderService) GetOrderList(ctx context.Context, req *service.OrderListRequest, resp *service.OrderListResponse) error {
	var count int64
	data := make([]*models.Order, 0)
	// 查询订单列表
	orders := models.DB.Model(new(models.Order))
	if req.UserId != 0 {
		orders = orders.Where("user_id = ?", req.UserId)
	}
	err := orders.Count(&count).Offset(int(req.Offset)).Limit(int(req.Limit)).Order("id desc").Find(&data).Error
	if err != nil {
		resp.Code = e.ERROR_DB_BASE
		return nil
	}
	resp.PageInfo = &service.PageInfo{
		Offset: req.Offset,
		Limit:  req.Limit,
		Total:  uint32(count),
	}
	resp.Code = e.SUCCESS
	resp.OrderList = schema.EncodeOrderList(data)

	return nil
}

// 微服务调用的 CreateOrder 方法
func (*OrderService) CreateOrder(ctx context.Context, req *service.OrderCreateRequest, resp *service.OrderCreateResponse) error {
	// 创建订单模型
	order := &models.Order{
		Name:   req.Name,
		UserID: req.UserId,
	}

	err := models.DB.Create(order).Error
	if err != nil {
		resp.Code = e.ERROR_DB_BASE
		return nil
	}

	resp.Code = e.SUCCESS
	resp.OrderDetail = schema.EncodeOrder(*order)
	return nil
}

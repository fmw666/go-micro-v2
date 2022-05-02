package core

import (
	"context"
	"fmt"
	"order/models"
	"order/service"
)

func BuildOrder(item models.Order) *service.OrderModel {
	orderModel := service.OrderModel{
		ID:     uint32(item.ID.ID),
		UserID: uint32(item.UserID),
	}
	return &orderModel
}

func (*OrderService) GetOrderList(ctx context.Context, req *service.OrderRequest, resp *service.OrderDetailResponse) error {
	var orderList []models.Order
	models.DB.Find(&orderList)
	resp.Code = 200
	// resp.OrderDetail = BuildOrder(orderList)
	return nil
}

func (*OrderService) GetOrderDetail(ctx context.Context, req *service.OrderRequest, resp *service.OrderDetailResponse) error {
	var order models.Order
	resp.Code = 200
	if err := models.DB.Where("id=?", req.ID).First(&order).Error; err != nil {
		fmt.Println("Order not found...")
		resp.Code = 400
		return nil
	}
	resp.OrderDetail = BuildOrder(order)
	return nil
}

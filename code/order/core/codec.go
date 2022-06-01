package core

import (
	"order/models"
	"order/service"
)

func buildOrder(item *models.Order) *service.OrderResponse {
	orderResp := service.OrderResponse{
		ID:     item.Id,
		Name:   item.Name,
		UserID: item.UserID,
	}
	return &orderResp
}

func buildOrderList(items []*models.Order) []*service.OrderResponse {
	orderList := make([]*service.OrderResponse, 0)
	for _, item := range items {
		orderList = append(orderList, buildOrder(item))
	}
	return orderList
}

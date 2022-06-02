package core

import (
	"order/models"
	"order/service"
)

func buildOrder(item *models.Order) *service.OrderResponse {
	orderResp := service.OrderResponse{
		Id:        item.Id,
		Name:      item.Name,
		UserId:    item.UserID,
		CreatedAt: item.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: item.UpdatedAt.Format("2006-01-02 15:04:05"),
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

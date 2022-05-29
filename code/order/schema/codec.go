package schema

import (
	"order/models"
	"order/service"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// 用于将模型 Order 转为 微服务响应结构
func EncodeOrder(item models.Order) *service.OrderResponse {
	return &service.OrderResponse{
		ID:     item.Id,
		Name:   item.Name,
		UserID: item.UserID,
		// Time 转 Timestamp
		CreatedAt: timestamppb.New(item.CreatedAt),
		UpdatedAt: timestamppb.New(item.UpdatedAt),
	}
}

func EncodeOrderList(orders []*models.Order) []*service.OrderResponse {
	var list []*service.OrderResponse
	for _, order := range orders {
		list = append(list, EncodeOrder(*order))
	}
	return list
}

// 用于将微服务响应结构 转为 schema 结构
func DecodeOrder(item *service.OrderResponse) *OrderResp {
	return &OrderResp{
		ID:     item.ID,
		Name:   item.Name,
		UserID: item.UserID,
		// Timestamp 转 Time
		CreatedAt: item.CreatedAt.AsTime().Format("2006-01-02 15:04:05"),
		UpdatedAt: item.UpdatedAt.AsTime().Format("2006-01-02 15:04:05"),
	}
}

func DecodeOrderList(orders []*service.OrderResponse) []*OrderResp {
	var list []*OrderResp
	for _, item := range orders {
		list = append(list, DecodeOrder(item))
	}
	return list
}

// 用于将微服务响应结构 转为 schema 结构
func DecodePageInfo(item *service.PageInfo) *PageInfoResp {
	return &PageInfoResp{
		Offset: item.Offset,
		Limit:  item.Limit,
		Total:  item.Total,
	}
}

package schema

import (
	"user/models"
	"user/service"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// 用于将模型 User 转为 微服务响应结构
func EncodeUser(item models.User) *service.UserModel {
	return &service.UserModel{
		ID:       item.Id,
		Username: item.Username,
		// Time 转 Timestamp
		CreatedAt: timestamppb.New(item.CreatedAt),
		UpdatedAt: timestamppb.New(item.UpdatedAt),
	}
}

// 用于将微服务响应结构 转为 schema 结构
func DecodeUser(item *service.UserModel) *UserResp {
	return &UserResp{
		ID:       item.ID,
		Username: item.Username,
		// Timestamp 转 Time
		CreatedAt: item.CreatedAt.AsTime().Format("2006-01-02 15:04:05"),
		UpdatedAt: item.UpdatedAt.AsTime().Format("2006-01-02 15:04:05"),
	}
}

// 用于将微服务响应结构 转为 schema 结构
func DecodePageInfo(item *service.PageInfo) *PageInfoResp {
	return &PageInfoResp{
		Offset: item.Offset,
		Limit:  item.Limit,
		Total:  item.Total,
	}
}

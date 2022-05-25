package schema

import (
	"user/models"
	"user/service"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func EncodeUser(item models.User) *service.UserModel {
	return &service.UserModel{
		ID:       item.Id,
		Username: item.Username,
		// Time 转 Timestamp
		CreatedAt: timestamppb.New(item.CreatedAt),
		UpdatedAt: timestamppb.New(item.UpdatedAt),
	}
}

func DecodeUser(item *service.UserModel) *UserResp {
	return &UserResp{
		ID:       item.ID,
		Username: item.Username,
		// Timestamp 转 Time
		CreatedAt: item.CreatedAt.AsTime().Format("2006-01-02 15:04:05"),
		UpdatedAt: item.UpdatedAt.AsTime().Format("2006-01-02 15:04:05"),
	}
}

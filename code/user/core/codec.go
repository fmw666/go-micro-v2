package core

import (
	"user/models"
	"user/service"
)

func buildUser(item models.User) *service.User {
	userResp := service.User{
		Id:        item.Id,
		Username:  item.Username,
		CreatedAt: item.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: item.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return &userResp
}

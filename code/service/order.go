package service

import (
	"app/config"
	"app/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetOrderList 获取订单列表
// @Summary 获取订单列表
// @Description 微服务模块 Order 中提供的获取订单列表服务
// @Tags Order 服务
// @Security ApiKeyAuth
// @Security BasicAuth
// @Accept  json
// @Produce  json
// @Param user_id query int false "用户ID"
// @Param offset query int false "偏移量"
// @Param limit query int false "限制数量"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /orders [get]
func GetOrderList(ginCtx *gin.Context) {
	offset, _ := strconv.Atoi(ginCtx.DefaultQuery("offset", config.AppSetting.DefaultOffset))
	limit, _ := strconv.Atoi(ginCtx.DefaultQuery("limit", config.AppSetting.DefaultLimit))
	userID, _ := strconv.Atoi(ginCtx.Query("user_id"))

	var count int64
	data := make([]*models.Order, 0)
	// 查询订单列表
	err := models.DB.Model(new(models.Order)).
		Where("user_id = ?", userID).
		Count(&count).Offset(offset).Limit(limit).Order("id desc").
		Find(&data).Error
	if err != nil {
		ginCtx.JSON(200, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}

	ginCtx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
		"page_info": gin.H{
			"total":  count,
			"offset": offset,
			"limit":  limit,
		},
	})
}

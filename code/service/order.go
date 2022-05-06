package service

import (
	"app/config"
	"app/models"
	"app/pkg/logger"
	"app/pkg/utils"
	"app/schema"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func BuildOrder(order models.Order) *schema.OrderResp {
	return &schema.OrderResp{
		ID:        uint(order.Id),
		Name:      order.Name,
		CreatedAt: order.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: order.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
}

// GetOrderList 获取订单列表
// @Summary 获取订单列表
// @Description Order 服务中提供的获取订单列表服务
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
	userID, _ := strconv.Atoi(ginCtx.DefaultQuery("user_id", "0"))

	var count int64
	data := make([]*models.Order, 0)
	// 查询订单列表
	orders := models.DB.Model(new(models.Order))
	if userID != 0 {
		orders = orders.Where("user_id = ?", userID)
	}
	err := orders.Count(&count).Offset(offset).Limit(limit).Order("id desc").Find(&data).Error
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

// CreateOrder 创建订单
// @Summary 创建订单
// @Description Order 服务中提供的创建订单服务
// @Tags Order 服务
// @Security ApiKeyAuth
// @Security BasicAuth
// @Accept  json
// @Produce  json
// @Param order body schema.OrderCreateReq true "订单"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /orders [post]
func CreateOrder(ginCtx *gin.Context) {
	var req schema.OrderCreateReq
	err := ginCtx.BindJSON(&req)
	if err != nil {
		ginCtx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}

	// 获取当前登录用户
	user := utils.GetCurrentUser(ginCtx)
	logger.Info("user: %+v", user)

	order := &models.Order{
		Name:   req.Name,
		UserID: uint(user.Id),
	}

	err = models.DB.Create(order).Error
	if err != nil {
		ginCtx.JSON(http.StatusOK, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}

	ginCtx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": BuildOrder(*order),
	})
}

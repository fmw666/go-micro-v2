package handlers

import (
	"api-gateway/service"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetOrderList 获取订单列表
// @Summary 获取订单列表
// @Description 微服务模块 Order 中提供的获取订单列表服务
// @Tags Order 服务
// @Accept  json
// @Produce  json
// @Param user_id query int false "用户 id"
// @Param offset query int false "偏移量"
// @Param limit query int false "限制数量"
// @Success 200 {string} json "{"code":0,"data":{},"message":""}"
// @Router /orders [get]
func GetOrderList(ginCtx *gin.Context) {
	// query 参数解析
	offset, _ := strconv.Atoi(ginCtx.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(ginCtx.DefaultQuery("limit", "10"))
	userID, _ := strconv.Atoi(ginCtx.DefaultQuery("user_id", "0"))

	// 构建 request
	req := service.OrderListRequest{
		UserId: uint32(userID),
		Offset: uint32(offset),
		Limit:  uint32(limit),
	}

	// 从 gin.Key 中取出服务实例
	orderService := ginCtx.Keys["orderService"].(service.OrderService)
	orderResp, err := orderService.GetOrderList(context.Background(), &req)
	PanicIfOrderError(err)

	ginCtx.JSON(http.StatusOK, orderResp)
}

// CreateOrder 创建订单
// @Summary 创建订单
// @Description 微服务模块 Order 中提供的创建订单服务
// @Tags Order 服务
// @Accept  json
// @Produce  json
// @Param order body schema.OrderCreateReq true "订单"
// @Success 200 {string} json "{"code":0,"data":{},"message":""}"
// @Router /orders [post]
func CreateOrder(ginCtx *gin.Context) {
	var orderReq service.OrderCreateRequest
	PanicIfOrderError(ginCtx.Bind(&orderReq))

	// 从 gin.Key 中取出服务实例
	orderService := ginCtx.Keys["orderService"].(service.OrderService)
	orderResp, err := orderService.CreateOrder(context.Background(), &orderReq)
	PanicIfOrderError(err)

	ginCtx.JSON(http.StatusOK, orderResp)
}

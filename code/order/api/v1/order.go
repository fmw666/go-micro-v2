package v1

import (
	"context"
	"order/config"
	"order/pkg/e"
	"order/pkg/utils"
	"order/schema"
	"order/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetOrderList 获取订单列表
// @Summary 获取订单列表
// @Description Order 服务中提供的获取订单列表服务
// @Tags Order 服务
// @Accept  json
// @Produce  json
// @Param user_id query int false "用户 id"
// @Param offset query int false "偏移量"
// @Param limit query int false "限制数量"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /orders [get]
func GetOrderList(ginCtx *gin.Context) {
	// query 参数解析
	offset, _ := strconv.Atoi(ginCtx.DefaultQuery("offset", config.AppSetting.DefaultOffset))
	limit, _ := strconv.Atoi(ginCtx.DefaultQuery("limit", config.AppSetting.DefaultLimit))
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
	if err != nil {
		utils.ErrorResponse(ginCtx, e.ERROR_SERVICE_BASE)
		return
	}
	if orderResp.Code != e.SUCCESS {
		utils.ErrorResponse(ginCtx, e.ErrorCode(orderResp.Code))
		return
	}
	utils.OkResponse(ginCtx, schema.DecodeOrderList(orderResp.OrderList), *schema.DecodePageInfo(orderResp.PageInfo))
}

// CreateOrder 创建订单
// @Summary 创建订单
// @Description Order 服务中提供的创建订单服务
// @Tags Order 服务
// @Accept  json
// @Produce  json
// @Param order body schema.OrderCreateReq true "订单"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /orders [post]
func CreateOrder(ginCtx *gin.Context) {
	// 获取 body 内容
	var req service.OrderCreateRequest
	if err := ginCtx.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(ginCtx, e.ERROR_PARAM_INVALID)
		return
	}
	// 从 gin.Key 中取出服务实例
	orderService := ginCtx.Keys["orderService"].(service.OrderService)
	orderResp, err := orderService.CreateOrder(context.Background(), &req)
	if err != nil {
		utils.ErrorResponse(ginCtx, e.ERROR_SERVICE_BASE)
		return
	}
	if orderResp.Code != e.SUCCESS {
		utils.ErrorResponse(ginCtx, e.ErrorCode(orderResp.Code))
		return
	}
	respData := schema.DecodeOrder(orderResp.OrderDetail)
	utils.OkResponse(ginCtx, respData)
}

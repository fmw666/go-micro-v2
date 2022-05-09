package v1

import (
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
// @Param user_id query int true "用户 id"
// @Param offset query int false "偏移量"
// @Param limit query int false "限制数量"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /orders [get]
func GetOrderList(ginCtx *gin.Context) {
	offset, _ := strconv.Atoi(ginCtx.DefaultQuery("offset", config.AppSetting.DefaultOffset))
	limit, _ := strconv.Atoi(ginCtx.DefaultQuery("limit", config.AppSetting.DefaultLimit))
	userID, _ := strconv.Atoi(ginCtx.DefaultQuery("user_id", "0"))

	var count int64
	var data interface{}
	var code e.ErrorCode
	switch {
	case userID > 0:
		count, data, code = service.GetOrderList(offset, limit, uint32(userID))
	case userID == 0:
		count, data, code = service.GetOrderList(offset, limit)
	}

	pageInfo := schema.PageInfoResp{
		Total:  count,
		Offset: int64(offset),
		Limit:  int64(limit),
	}
	utils.Response(ginCtx, code, data, pageInfo)
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
	var req schema.OrderCreateReq
	err := ginCtx.BindJSON(&req)
	if err != nil {
		utils.ErrorResponse(ginCtx, e.ERROR_PARAM_INVALID)
		return
	}
	data, code := service.CreateOrder(req.Name, req.UserID)
	utils.Response(ginCtx, code, data)
}

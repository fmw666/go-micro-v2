package v1

import (
	"app/config"
	"app/schema"
	"app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetOrderList 获取订单列表
// @Summary 获取订单列表
// @Description Order 服务中提供的获取订单列表服务
// @Tags Order 服务
// @Accept json
// @Produce json
// @Param user_id query int false "用户 id"
// @Param offset query int false "偏移量"
// @Param limit query int false "限制数量"
// @Success 200 {object} schema.Response "{"code":0,"data":{},"message":""}"
// @Router /orders [get]
func GetOrderList(ginCtx *gin.Context) {
	offset, _ := strconv.Atoi(ginCtx.DefaultQuery("offset", config.AppSetting.DefaultOffset))
	limit, _ := strconv.Atoi(ginCtx.DefaultQuery("limit", config.AppSetting.DefaultLimit))
	userID, _ := strconv.Atoi(ginCtx.DefaultQuery("user_id", "0"))

	ginCtx.JSON(http.StatusOK, service.GetOrderList(uint32(offset), uint32(limit), uint32(userID)))
}

// CreateOrder 创建订单
// @Summary 创建订单
// @Description Order 服务中提供的创建订单服务
// @Tags Order 服务
// @Accept json
// @Produce json
// @Param order body schema.OrderCreateReq true "订单"
// @Success 200 {object} schema.Response "{"code":0,"data":{},"message":""}"
// @Router /orders [post]
func CreateOrder(ginCtx *gin.Context) {
	var req schema.OrderCreateReq
	if err := ginCtx.BindJSON(&req); err != nil {
		panic(err)
	}

	ginCtx.JSON(http.StatusOK, service.CreateOrder(req.Name, req.UserID))
}

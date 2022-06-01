package handlers

import "github.com/gin-gonic/gin"

// GetOrderList 获取订单列表
// @Summary 获取订单列表
// @Description 微服务模块 Order 中提供的获取订单列表服务
// @Tags Order 服务
// @Accept  json
// @Produce  json
// @Param user_id query int false "用户 id"
// @Param offset query int false "偏移量"
// @Param limit query int false "限制数量"
// @Success 200 {string} json "{"code":200,"data":{},"msg":""}"
// @Router /orders [get]
func GetOrderList(ginCtx *gin.Context) {

}

// CreateOrder 创建订单
// @Summary 创建订单
// @Description 微服务模块 Order 中提供的创建订单服务
// @Tags Order 服务
// @Accept  json
// @Produce  json
// @Param order body schema.OrderCreateReq true "订单"
// @Success 200 {string} json "{"code":200,"data":{},"msg":""}"
// @Router /orders [post]
func CreateOrder(ginCtx *gin.Context) {

}

package handlers

import "github.com/gin-gonic/gin"

// GetOrderList 获取订单列表
// @Summary 获取订单列表
// @Description 微服务模块 Order 中提供的获取订单列表服务
// @Tags Order 服务
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /orders [get]
func GetOrderList(ginCtx *gin.Context) {

}

// GetOrderDetail 获取订单详情
// @Summary 获取订单详情
// @Description 微服务模块 Order 中提供的获取订单详情服务
// @Tags Order 服务
// @Security ApiKeyAuth
// @Accept  json
// @Produce  json
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /orders/:id [get]
func GetOrderDetail(ginCtx *gin.Context) {

}

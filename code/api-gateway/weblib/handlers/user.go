package handlers

import (
	"api-gateway/service"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册
// @Summary 用户注册
// @Description 微服务模块 User 中提供的用户注册服务
// @Tags User 服务
// @Accept  json
// @Produce  json
// @Param body body service.UserRegisterRequest true "注册"
// @Success 200 {string} json "{"code":200,"data":{},"message":""}"
// @Router /user/register [post]
func UserRegister(ginCtx *gin.Context) {
	var userReq service.UserRegisterRequest
	PanicIfUserError(ginCtx.Bind(&userReq))
	// 从 gin.Key 中取出服务实例
	userService := ginCtx.Keys["userService"].(service.UserService)
	userResp, err := userService.UserRegister(context.Background(), &userReq)
	PanicIfUserError(err)

	ginCtx.JSON(http.StatusOK, userResp)
}

// UserLogin 用户登录
// @Summary 用户登录
// @Description 微服务模块 User 中提供的用户登录服务
// @Tags User 服务
// @Accept  json
// @Produce  json
// @Param body body service.UserLoginRequest true "登录"
// @Success 200 {string} json "{"code":200,"data":{},"message":""}"
// @Router /user/login [post]
func UserLogin(ginCtx *gin.Context) {
	var userReq service.UserLoginRequest
	PanicIfUserError(ginCtx.Bind(&userReq))
	// 从 gin.Key 中取出服务实例
	userService := ginCtx.Keys["userService"].(service.UserService)
	userResp, err := userService.UserLogin(context.Background(), &userReq)
	PanicIfUserError(err)

	ginCtx.JSON(http.StatusOK, userResp)
}

// UserOrderCreate 用户创建订单
// @Summary 用户创建订单
// @Description User 服务中提供的用户创建订单服务
// @Tags User 服务
// @Security ApiKeyAuth
// @Security BasicAuth
// @Accept  json
// @Produce  json
// @Param body body service.OrderCreateRequest true "订单"
// @Success 200 {string} json "{"code":0,"data":{}}"
// @Router /user/orders [post]
func UserOrderCreate(ginCtx *gin.Context) {
	var orderReq service.OrderCreateRequest
	PanicIfUserError(ginCtx.Bind(&orderReq))

	// 获取当前登录用户
	user_id := ginCtx.Keys["user_id"].(uint32)
	orderReq.UserId = user_id

	// 从 gin.Key 中取出服务实例
	orderService := ginCtx.Keys["orderService"].(service.OrderService)
	orderResp, err := orderService.CreateOrder(context.Background(), &orderReq)
	PanicIfUserError(err)

	ginCtx.JSON(http.StatusOK, orderResp)
}

// GetUserOrderList 用户订单列表
// @Summary 用户订单列表
// @Description User 服务中提供的用户订单列表服务
// @Tags User 服务
// @Security ApiKeyAuth
// @Security BasicAuth
// @Accept  json
// @Produce  json
// @Param offset query int false "Offset"
// @Param limit query int false "Limit"
// @Success 200 {string} json "{"code":0,"data":{}}"
// @Router /user/orders [get]
func GetUserOrderList(ginCtx *gin.Context) {
	// query 参数解析
	offset, _ := strconv.Atoi(ginCtx.DefaultQuery("offset", "0"))
	limit, _ := strconv.Atoi(ginCtx.DefaultQuery("limit", "10"))

	// 构建 req
	orderReq := service.OrderListRequest{
		Offset: uint32(offset),
		Limit:  uint32(limit),
	}

	// 获取当前登录用户
	user_id := ginCtx.Keys["user_id"].(uint32)
	orderReq.UserId = user_id

	// 从 gin.Key 中取出服务实例
	orderService := ginCtx.Keys["orderService"].(service.OrderService)
	orderResp, err := orderService.GetOrderList(context.Background(), &orderReq)
	PanicIfUserError(err)

	ginCtx.JSON(http.StatusOK, orderResp)
}

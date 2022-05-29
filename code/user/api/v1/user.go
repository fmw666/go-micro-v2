package v1

import (
	"context"
	"strconv"
	"user/config"
	"user/models"
	"user/pkg/e"
	"user/pkg/utils"
	"user/schema"
	"user/service"

	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册
// @Summary 用户注册
// @Description User 服务中提供的用户注册服务
// @Tags User 服务
// @Accept  json
// @Produce  json
// @Param body body schema.RegisterReq true "注册"
// @Success 200 {string} json "{"code":0,"data":{}}"
// @Router /user/register [post]
func UserRegister(ginCtx *gin.Context) {
	// 获取 body 内容
	var req service.UserRegisterRequest
	if err := ginCtx.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(ginCtx, e.ERROR_PARAM_INVALID)
		return
	}
	// 从 gin.Key 中取出服务实例
	userService := ginCtx.Keys["userService"].(service.UserService)
	userResp, err := userService.UserRegister(context.Background(), &req)
	if err != nil {
		utils.ErrorResponse(ginCtx, e.ERROR_SERVICE_BASE)
		return
	}
	if userResp.Code != e.SUCCESS {
		utils.ErrorResponse(ginCtx, e.ErrorCode(userResp.Code))
		return
	}
	token, _ := utils.GenerateToken(uint(userResp.UserDetail.ID))
	respData := gin.H{
		"token": token,
		"user":  schema.DecodeUser(userResp.UserDetail),
	}
	utils.OkResponse(ginCtx, respData)
}

// UserLogin 用户登录
// @Summary 用户登录
// @Description User 服务中提供的用户登录服务
// @Tags User 服务
// @Accept  json
// @Produce  json
// @Param body body schema.LoginReq true "登录"
// @Success 200 {string} json "{"code":0,"data":{}}"
// @Router /user/login [post]
func UserLogin(ginCtx *gin.Context) {
	// 获取 body 内容
	var req service.UserLoginRequest
	if err := ginCtx.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(ginCtx, e.ERROR_PARAM_INVALID)
		return
	}
	// 从 gin.Key 中取出服务实例
	userService := ginCtx.Keys["userService"].(service.UserService)
	userResp, err := userService.UserLogin(context.Background(), &req)
	if err != nil {
		utils.ErrorResponse(ginCtx, e.ERROR_SERVICE_BASE)
		return
	}
	if userResp.Code != e.SUCCESS {
		utils.ErrorResponse(ginCtx, e.ErrorCode(userResp.Code))
		return
	}
	token, _ := utils.GenerateToken(uint(userResp.UserDetail.ID))
	respData := gin.H{
		"token": token,
		"user":  schema.DecodeUser(userResp.UserDetail),
	}
	utils.OkResponse(ginCtx, respData)
}

// UserOrderCreate 用户创建订单
// @Summary 用户创建订单
// @Description User 服务中提供的用户创建订单服务
// @Tags User 服务
// @Security ApiKeyAuth
// @Security BasicAuth
// @Accept  json
// @Produce  json
// @Param body body schema.UserOrderCreateReq true "订单"
// @Success 200 {string} json "{"code":0,"data":{}}"
// @Router /user/orders [post]
func UserOrderCreate(ginCtx *gin.Context) {
	// 获取 body 内容
	var req service.OrderCreateRequest
	if err := ginCtx.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(ginCtx, e.ERROR_PARAM_INVALID)
		return
	}

	// 获取当前登录用户
	user := ginCtx.Keys["user"].(models.User)
	req.UserId = user.Id

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
	utils.OkResponse(ginCtx, schema.DecodeOrder(orderResp.OrderDetail))
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
	offset, _ := strconv.Atoi(ginCtx.DefaultQuery("offset", config.AppSetting.DefaultOffset))
	limit, _ := strconv.Atoi(ginCtx.DefaultQuery("limit", config.AppSetting.DefaultLimit))

	// 获取当前登录用户
	user := ginCtx.Keys["user"].(models.User)

	// 构建 request
	req := service.OrderListRequest{
		UserId: user.Id,
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

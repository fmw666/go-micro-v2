package v1

import (
	"bytes"
	"net/http"
	"user/pkg/e"
	"user/pkg/utils"
	"user/pkg/utils/consul"
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
	var req schema.RegisterReq
	if err := ginCtx.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(ginCtx, e.ERROR_PARAM_INVALID)
		return
	}
	data, code := service.UserRegister(req.Username, req.Password, req.PasswordConfirm)
	utils.Response(ginCtx, code, data)
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
	var req schema.LoginReq
	if err := ginCtx.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponse(ginCtx, e.ERROR_PARAM_INVALID)
		return
	}
	data, code := service.UserLogin(req.Username, req.Password)
	utils.Response(ginCtx, code, data)
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
	var req schema.UserOrderCreateReq
	err := ginCtx.BindJSON(&req)
	if err != nil {
		utils.ErrorResponse(ginCtx, e.ERROR_PARAM_INVALID)
		return
	}
	// 获取 Order 服务地址
	hostAddress, err := consul.GetServiceAddr("rpcOrderService")
	if err != nil || hostAddress == "" {
		utils.ErrorResponse(ginCtx, e.ERROR_SERVICE_NOT_FOUND)
		return
	}
	// 调用 Order 服务
	url := "http://" + hostAddress + "/api/v1/orders"
	resp, _ := http.Post(url, "application/json;charset=utf-8", bytes.NewBuffer([]byte("")))
	// 获取当前登录用户
	// user := ginCtx.Keys["user"].(models.User)
	// 	data, code := service.CreateOrder(req.Name, user.Id)
	// 	utils.Response(ginCtx, code, data)
	ginCtx.JSON(200, gin.H{
		"code": 0,
		"data": resp,
	})
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
	// offset, _ := strconv.Atoi(ginCtx.DefaultQuery("offset", config.AppSetting.DefaultOffset))
	// limit, _ := strconv.Atoi(ginCtx.DefaultQuery("limit", config.AppSetting.DefaultLimit))

	// 获取 Order 服务地址
	hostAddress, err := consul.GetServiceAddr("rpcOrderService")
	if err != nil || hostAddress == "" {
		utils.ErrorResponse(ginCtx, e.ERROR_SERVICE_NOT_FOUND)
		return
	}
	// 调用 Order 服务
	url := "http://" + hostAddress + "/api/v1/orders"
	resp, _ := http.Post(url, "application/json;charset=utf-8", bytes.NewBuffer([]byte("")))
	ginCtx.JSON(200, gin.H{
		"code": 0,
		"data": resp,
	})

	// // 获取当前登录用户
	// user := ginCtx.Keys["user"].(models.User)
	// count, data, code := service.GetOrderList(offset, limit, user.Id)
	// pageInfo := schema.PageInfoResp{
	// 	Total:  count,
	// 	Offset: int64(offset),
	// 	Limit:  int64(limit),
	// }
	// utils.Response(ginCtx, code, data, pageInfo)
}

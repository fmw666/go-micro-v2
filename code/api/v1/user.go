package v1

import (
	"app/config"
	"app/models"
	"app/schema"
	"app/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserRegister 用户注册
// @Summary 用户注册
// @Description User 服务中提供的用户注册服务
// @Tags User 服务
// @Accept json
// @Produce json
// @Param body body schema.RegisterReq true "注册"
// @Success 200 {object} schema.Response "{"code":0,"data":{},"message":""}"
// @Router /user/register [post]
func UserRegister(ginCtx *gin.Context) {
	// 获取 body 内容
	var req schema.RegisterReq
	if err := ginCtx.ShouldBindJSON(&req); err != nil {
		panic(err)
	}

	ginCtx.JSON(http.StatusOK, service.UserRegister(req.Username, req.Password, req.PasswordConfirm))
}

// UserLogin 用户登录
// @Summary 用户登录
// @Description User 服务中提供的用户登录服务
// @Tags User 服务
// @Accept json
// @Produce json
// @Param body body schema.LoginReq true "登录"
// @Success 200 {object} schema.Response "{"code":0,"data":{},"message":""}"
// @Router /user/login [post]
func UserLogin(ginCtx *gin.Context) {
	// 获取 body 内容
	var req schema.LoginReq
	if err := ginCtx.ShouldBindJSON(&req); err != nil {
		panic(err)
	}

	ginCtx.JSON(http.StatusOK, service.UserLogin(req.Username, req.Password))
}

// UserOrderCreate 用户创建订单
// @Summary 用户创建订单
// @Description User 服务中提供的用户创建订单服务
// @Tags User 服务
// @Security ApiKeyAuth
// @Security BasicAuth
// @Accept json
// @Produce json
// @Param body body schema.UserOrderCreateReq true "订单"
// @Success 200 {object} schema.Response "{"code":0,"data":{},"message":""}"
// @Router /user/orders [post]
func UserOrderCreate(ginCtx *gin.Context) {
	var req schema.UserOrderCreateReq
	err := ginCtx.BindJSON(&req)
	if err != nil {
		panic(err)
	}
	// 获取当前登录用户
	user := ginCtx.Keys["user"].(models.User)

	ginCtx.JSON(http.StatusOK, service.CreateOrder(req.Name, user.Id))
}

// GetUserOrderList 用户订单列表
// @Summary 用户订单列表
// @Description User 服务中提供的用户订单列表服务
// @Tags User 服务
// @Security ApiKeyAuth
// @Security BasicAuth
// @Accept json
// @Produce json
// @Param offset query int false "Offset"
// @Param limit query int false "Limit"
// @Success 200 {object} schema.Response "{"code":0,"data":{},"message":""}"
// @Router /user/orders [get]
func GetUserOrderList(ginCtx *gin.Context) {
	offset, _ := strconv.Atoi(ginCtx.DefaultQuery("offset", config.AppSetting.DefaultOffset))
	limit, _ := strconv.Atoi(ginCtx.DefaultQuery("limit", config.AppSetting.DefaultLimit))

	// 获取当前登录用户
	user := ginCtx.Keys["user"].(models.User)

	ginCtx.JSON(http.StatusOK, service.GetOrderList(uint32(offset), uint32(limit), user.Id))
}

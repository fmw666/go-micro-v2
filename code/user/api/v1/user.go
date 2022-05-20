package v1

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"user/config"
	"user/models"
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
	utils.Response(ginCtx, e.SUCCESS, userResp)
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
	utils.Response(ginCtx, e.SUCCESS, userResp)
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
	// 获取当前登录用户
	user := ginCtx.Keys["user"].(models.User)
	// 调用 Order 服务
	url := "http://" + hostAddress + "/api/v1/orders"
	body := bytes.NewBuffer([]byte("{\"name\":\"" + req.Name + "\",\"user_id\":" + strconv.FormatInt(int64(user.Id), 10) + "}"))
	resp, _ := http.Post(url, "application/json;charset=utf-8", body)

	respData, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	var data interface{}
	json.Unmarshal(respData, &data)

	ginCtx.JSON(http.StatusOK, data)
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
	offset := ginCtx.DefaultQuery("offset", config.AppSetting.DefaultOffset)
	limit := ginCtx.DefaultQuery("limit", config.AppSetting.DefaultLimit)

	// 获取 Order 服务地址
	hostAddress, err := consul.GetServiceAddr(config.ServiceSetting.OrderServiceName)
	if err != nil || hostAddress == "" {
		utils.ErrorResponse(ginCtx, e.ERROR_SERVICE_NOT_FOUND)
		return
	}
	// 获取当前登录用户
	user := ginCtx.Keys["user"].(models.User)
	// 调用 Order 服务
	url := "http://" + hostAddress + "/api/v1/orders"
	url += "?offset=" + offset + "&limit=" + limit + "&user_id=" + strconv.FormatInt(int64(user.Id), 10)
	resp, _ := http.Get(url)

	respData, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	var data interface{}
	json.Unmarshal(respData, &data)

	ginCtx.JSON(http.StatusOK, data)
}

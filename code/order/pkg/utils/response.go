package utils

import (
	"net/http"
	"order/pkg/e"
	"order/schema"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(ginCtx *gin.Context, code e.ErrorCode) {
	Response(ginCtx, code, nil)
}

func OkResponse(ginCtx *gin.Context, data any, pageInfo ...schema.PageInfoResp) {
	Response(ginCtx, e.SUCCESS, data, pageInfo...)
}

func Response(ginCtx *gin.Context, code e.ErrorCode, data any, pageInfo ...schema.PageInfoResp) {
	if code != e.SUCCESS {
		ginCtx.JSON(http.StatusOK, gin.H{"code": code, "msg": e.GetMsg(code)})
		return
	}
	if len(pageInfo) > 0 {
		ginCtx.JSON(http.StatusOK, gin.H{"code": code, "data": data, "page_info": pageInfo[0]})
		return
	}
	ginCtx.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "data": data})
}

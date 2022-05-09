package utils

import (
	"net/http"
	"user/pkg/e"
	"user/schema"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(ginCtx *gin.Context, code e.ErrorCode) {
	ginCtx.JSON(http.StatusOK, gin.H{"code": code, "msg": e.GetMsg(code)})
}

func Response(ginCtx *gin.Context, code e.ErrorCode, data interface{}, pageInfo ...schema.PageInfoResp) {
	if code != e.SUCCESS {
		ErrorResponse(ginCtx, code)
		return
	}
	if len(pageInfo) > 0 {
		ginCtx.JSON(http.StatusOK, gin.H{"code": code, "data": data, "page_info": pageInfo[0]})
		return
	}
	ginCtx.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "data": data})
}

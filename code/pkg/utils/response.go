package utils

import (
	"app/pkg/e"
	"app/schema"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(ginCtx *gin.Context, code e.ErrorCode, data interface{}, page_info ...schema.PageInfo) {
	if code != e.SUCCESS {
		ginCtx.JSON(http.StatusOK, gin.H{"code": code, "msg": e.GetMsg(code)})
		return
	}
	if len(page_info) > 0 {
		ginCtx.JSON(http.StatusOK, gin.H{"code": code, "data": data, "page_info": page_info[0]})
		return
	}
	ginCtx.JSON(http.StatusOK, gin.H{"code": e.SUCCESS, "data": data})
}

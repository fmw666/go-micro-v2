package schema

import (
	"app/pkg/e"
)

type BaseResponse struct {
	Code    e.ErrorCode `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    any         `json:"data"`
}

type Response struct {
	BaseResponse
}

type ResponseWithPageInfo struct {
	BaseResponse
	PageInfo *PageInfo `json:"page_info,omitempty"`
}

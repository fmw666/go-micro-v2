package schema

import (
	"app/pkg/e"
)

type Response struct {
	Code     e.ErrorCode `json:"code"`
	Message  string      `json:"message,omitempty"`
	Data     any         `json:"data"`
	PageInfo *PageInfo   `json:"page_info,omitempty"`
}

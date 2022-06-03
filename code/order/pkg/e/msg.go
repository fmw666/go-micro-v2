package e

// GetMsg 获取错误信息
func GetMsg(code ErrorCode) string {
	if msg, ok := MsgMaps[code]; ok {
		return msg
	}
	return MsgMaps[ERROR_BASE]
}

var MsgMaps = map[ErrorCode]string{
	// 基础错误
	ERROR_BASE: "未知错误",

	// 数据库错误 11000-11999
	ERROR_DB_BASE:   "数据库错误",
	ERROR_DB_CREATE: "创建数据失败",

	// 鉴权错误 12000-12999
	ERROR_AUTH_BASE: "鉴权错误",
	ERROR_AUTH_FAIL: "鉴权失败",

	// 参数错误 13000-13999
	ERROR_PARAM_BASE:        "参数错误",
	ERROR_PARAM_INVALID:     "参数无效",
	ERROR_PARAM_NOT_CONTENT: "参数不能为空",

	// 服务错误 14000-14999
	ERROR_SERVICE_BASE:      "服务错误",
	ERROR_SERVICE_NOT_FOUND: "服务未找到",

	// 用户类错误 15000 ~ 15999
	ERROR_USER_BASE:          "用户错误",
	ERROR_USER_EXIST:         "用户名已存在",
	ERROR_USER_NOT_FOUND:     "用户名不存在",
	ERROR_USER_PASSWORD:      "密码错误",
	ERROR_USER_SET_PASSWORD:  "设置密码失败",
	ERROR_PASSWORD_NOT_MATCH: "密码不匹配",

	// MQ 错误 20000 ~ 21000
	ERROR_MQ_BASE:    "MQ 服务错误",
	ERROR_MQ_PUBLISH: "MQ 发布消息失败",
}

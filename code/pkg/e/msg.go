package e

var MsgMaps = map[ErrorCode]string{
	ERROR_BASE:               "未知错误",
	ERROR_DB_BASE:            "数据库错误",
	ERROR_DB_CREATE:          "创建数据失败",
	ERROR_AUTH_BASE:          "鉴权错误",
	ERROR_AUTH_FAIL:          "鉴权失败",
	ERROR_PARAM_BASE:         "参数错误",
	ERROR_PARAM_INVALID:      "参数无效",
	ERROR_USER_BASE:          "用户错误",
	ERROR_USER_EXIST:         "用户名已存在",
	ERROR_USER_NOT_FOUND:     "用户名不存在",
	ERROR_USER_PASSWORD:      "密码错误",
	ERROR_PASSWORD_NOT_MATCH: "密码不匹配",
}

// GetMsg 获取错误信息
func GetMsg(code ErrorCode) string {
	if msg, ok := MsgMaps[code]; ok {
		return msg
	}
	return MsgMaps[ERROR_BASE]
}

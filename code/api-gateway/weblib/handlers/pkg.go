package handlers

import (
	"api-gateway/pkg/logger"
	"errors"
)

// 包装错误
func PanicIfUserError(err error) {
	if err != nil {
		err = errors.New("userService--" + err.Error())
		logger.Info(err)
		panic(err)
	}
}

func PanicIfOrderError(err error) {
	if err != nil {
		err = errors.New("orderService--" + err.Error())
		logger.Info(err)
		panic(err)
	}
}

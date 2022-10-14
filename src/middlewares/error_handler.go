package middlewares

import (
	"attendance_backend/src/constant"
	"attendance_backend/src/utils"
	"github.com/kataras/iris/v12"
)

// ErrorBefore 前置处理
func ErrorBefore(ctx iris.Context) {
	ctx.Values().Set(constant.ErrorList, []error{})
	ctx.Values().Set(constant.ErrorCode, constant.Success)
	ctx.Values().Set(constant.ErrorMsg, "success")
	ctx.Next()
}

// ErrorAfter 后置处理
func ErrorAfter(ctx iris.Context) {
	errorCode := ctx.Values().Get(constant.ErrorCode)
	if errorCode != constant.Success {
		ctx.Next()
		return
	}
	errorMsg := ctx.Values().Get(constant.ErrorMsg)
	errorListInterface := ctx.Values().Get(constant.ErrorList)
	errorList := errorListInterface.([]error)
	var data []string
	for _, err := range errorList {
		data = append(data, err.Error())
	}
	utils.ReturnJson(ctx, errorCode.(int), errorMsg.(string), data)
	ctx.Next()
}

package middlewares

import (
	"attendance_backend/src/constant"
	"github.com/kataras/iris/v12"
)

// ErrorBefore 前置处理
func ErrorBefore(ctx iris.Context) {
	ctx.Values().Set(constant.ErrorList, []error{})
	ctx.Values().Set(constant.ErrorCode, constant.Success)
	ctx.Values().Set(constant.ErrorMsg, "success")
	ctx.Next()
}

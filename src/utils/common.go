package utils

import (
	"attendance_backend/src/constant"
	"attendance_backend/src/entities/_interface"
	"attendance_backend/src/utils/log"
	"github.com/kataras/iris/v12"
	"net/http"
)

func ReturnJson(ctx iris.Context, ret int, msg string, data ...interface{}) {
	err := ctx.StopWithJSON(http.StatusOK, _interface.ResponseStruct{
		Ret:  ret,
		Msg:  msg,
		Data: data,
	})
	if err != nil {
		log.Logger.Errorf("[return json] stop with json failed,%v", err)
	}
}

func SetError(ctx iris.Context, ret int, msg string, err error) {
	ctx.Values().Set(constant.ErrorCode, ret)
	ctx.Values().Set(constant.ErrorMsg, msg)
	errList := ctx.Values().Get(constant.ErrorList)
	errList = append(errList.([]error), err)
}

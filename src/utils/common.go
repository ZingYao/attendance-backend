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

func ReturnErrJson(ctx iris.Context, ret int, msg string) {
	errorListInterface := ctx.Values().Get(constant.ErrorList)
	errorList := errorListInterface.([]error)
	var data []string
	for _, err := range errorList {
		data = append(data, err.Error())
	}
	ReturnJson(ctx, ret, msg, data)
}

func SetError(ctx iris.Context, err error) {
	errList := ctx.Values().Get(constant.ErrorList)
	errList = append(errList.([]error), err)
	ctx.Values().Set(constant.ErrorList, errList)
}

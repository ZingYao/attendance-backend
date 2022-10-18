package v1

import (
	"attendance_backend/src/constant"
	"attendance_backend/src/repository"
	"attendance_backend/src/service"
	"attendance_backend/src/utils"
	"fmt"
	"github.com/kataras/iris/v12"
)

type (
	IndexController interface {
		Index(ctx iris.Context)
		Error(ctx iris.Context)
	}
	indexController struct {
		indexService service.IndexService
	}
)

func (c indexController) Index(ctx iris.Context) {
	utils.ReturnJson(ctx, constant.Success, "this is api root")
}

func (c indexController) Error(ctx iris.Context) {
	utils.SetError(ctx, fmt.Errorf("test1"))
	utils.SetError(ctx, fmt.Errorf("test2"))
	utils.SetError(ctx, fmt.Errorf("test3"))
	utils.ReturnErrJson(ctx, constant.Failed, "failed")
}

func GetIndexController(repo repository.Factory) IndexController {
	return &indexController{
		indexService: service.GetIndexService(repo),
	}
}

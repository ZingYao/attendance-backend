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
	}
	indexController struct {
		indexService service.IndexService
	}
)

func (c indexController) Index(ctx iris.Context) {
	utils.ReturnJson(ctx, constant.Success, "this is api root")
	utils.SetError(ctx, constant.Failed, "test", fmt.Errorf("hello"))
	ctx.Next()
}

func GetIndexController(repo repository.Factory) IndexController {
	return &indexController{
		indexService: service.GetIndexService(repo),
	}
}

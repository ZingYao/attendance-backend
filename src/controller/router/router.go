package router

import (
	v1 "attendance_backend/src/controller/v1"
	"attendance_backend/src/repository/postgre_sql"
	"github.com/kataras/iris/v12"
)

func Register(api iris.Party) {
	{
		//Index 注册
		repo := postgre_sql.GetRepository()
		indexController := v1.GetIndexController(repo)
		api.Get("/", indexController.Index)
		api.Get("/err", indexController.Error)
	}
}

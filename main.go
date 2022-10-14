package main

import (
	"attendance_backend/src/controller/router"
	"attendance_backend/src/dbsource"
	"attendance_backend/src/middlewares"
	"attendance_backend/src/utils/config"
	_ "attendance_backend/src/utils/config"
	"attendance_backend/src/utils/log"
	"fmt"
	"github.com/kataras/iris/v12"
	"io"
	"os"
	"time"
)

func main() {
	conf := config.GetConfig()
	var err error
	app := iris.New()

	//日志相关内容配置
	{
		//日志相关配置
		app.Logger().SetLevel(conf.Log.Level)
		var output io.Writer
		switch conf.Log.Output {
		case "console":
			output = os.Stdout
		case "file":
			err = os.Mkdir(conf.Log.FilePath, os.FileMode(0755))
			if err != nil {
				app.Logger().Warnf("create log dir failed,%v", err)
			}
			filePath := fmt.Sprintf("%s%c%s.log", conf.Log.FilePath, os.PathSeparator, time.Now().Format("2006_01_02"))
			output, err = os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.FileMode(0655))
			if err != nil {
				panic(fmt.Errorf("create config file failed,%v", err))
			}
		default:
			panic(fmt.Errorf("config log.output not support(it's must in [console,file])"))
		}
		app.Logger().SetOutput(output)
		log.Logger = app.Logger()
	}

	//数据库初始化
	dbsource.PostgreSql{}.InitPostgreSql(conf.Database.PostgreSql)
	dbsource.Redis{}.InitRedis(conf.Database.Redis)

	//api初始化
	api := app.Party("/api", middlewares.IpInterceptor, middlewares.ErrorBefore)
	router.Register(api)
	api.Done(middlewares.ErrorAfter)
	err = app.Listen(
		conf.App.Addr, //绑定地址
		//iris.WithTimeout(time.Duration(conf.App.Timeout)*time.Millisecond), //指定超时时间
		iris.WithLogLevel(conf.Log.Level),                      //指定日志等级
		iris.WithCharset("utf-8"),                              // 指定编码
		iris.WithPostMaxMemory(conf.App.MaxPostMemory*iris.MB), //指定post数据最大大小
		iris.WithTimeFormat("2006-01-02 15:04:05"),             //指定时间格式
	)
	if err != nil {
		app.Logger().Errorf("run application:%s failed,%v", conf.App.Name, err)
	}
}

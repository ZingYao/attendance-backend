package middlewares

import (
	"attendance_backend/src/constant"
	"attendance_backend/src/dbsource"
	"attendance_backend/src/utils"
	"attendance_backend/src/utils/config"
	"attendance_backend/src/utils/log"
	"fmt"
	"github.com/kataras/iris/v12"
	"time"
)

// IpInterceptor 利用redis限制单IP访问的qps
func IpInterceptor(ctx iris.Context) {
	conf := config.GetConfig()
	key := fmt.Sprintf(constant.IpInterceptorKey, ctx.RemoteAddr())
	//获取redis值
	strCmd := dbsource.GetRedisConn(constant.IpInterceptorDb).Get(ctx.Request().Context(), key)
	//获取出错当0处理
	times, err := strCmd.Int()
	if err != nil && err.Error() != "redis: nil" {
		log.Logger.Warnf("[ip interceptor]redis trans times to int failed,%v", err)
	}
	//判断是否有过多的请求
	if times >= conf.App.Qps {
		utils.ReturnJson(ctx, constant.TooManyRequest, "too many request")
		return
	}
	if times == 0 {
		//设置redis key过期时间一秒，每多请求一次时间重置一次
		boolCmd := dbsource.GetRedisConn(constant.IpInterceptorDb).SetEX(ctx, key, 1, 1*time.Second)
		err = boolCmd.Err()
		if err != nil {
			log.Logger.Warnf("[ip interceptor]redis set %q nx failed,%v", key, err)
			dbsource.GetRedisConn(constant.IpInterceptorDb).Del(ctx, key)
			//设置自动删除时间失败则直接删除key 避免影响后续请求
		}
	} else {
		//设置请求数+1
		intCmd := dbsource.GetRedisConn(constant.IpInterceptorDb).Incr(ctx, key)
		err = intCmd.Err()
		if err != nil {
			log.Logger.Warnf("[ip interceptor]redis incr failed,%v", err)
		}
	}
	ctx.Next()
}

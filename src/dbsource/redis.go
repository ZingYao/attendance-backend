package dbsource

import (
	"attendance_backend/src/entities"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type Redis struct {
}

var redisConf = redis.Options{}

var redisConn *redis.Client

func (Redis) InitRedis(redisConfig entities.RedisStruct) {
	redisConf.Addr = fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port)
	redisConf.Password = redisConfig.Pass
	redisConf.DB = 0
	redisConn = redis.NewClient(&redisConf)
}

func GetRedisConn(db int) *redis.Client {
	if db >= 16 || db < 0 {
		return nil
	}
	redisConf.DB = db
	return redisConn
}

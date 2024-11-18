package drivers

import (
	"strconv"
	"sync"
	"wolf/config"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client = nil
var redisOnce sync.Once

func GetRedisConnection() *redis.Client {
	redisOnce.Do(func() {
		conf := config.GetDeployConfig().Redis
		rdb = redis.NewClient(&redis.Options{
			Addr:     conf.Host + strconv.Itoa(conf.Port),
			Password: conf.Password,
			DB:       conf.Database,
		})
	})

	return rdb
}

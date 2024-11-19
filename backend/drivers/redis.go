package drivers

import (
	"context"
	"strconv"
	"time"
	"wolf/config"

	"github.com/redis/go-redis/v9"
)

var Redis *redis.Client

func init() {
	conf := config.GetDeployConfig().Redis
	Redis = redis.NewClient(&redis.Options{
		Addr:     conf.Host + ":" + strconv.Itoa(conf.Port),
		Password: conf.Password,
		DB:       conf.Database,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	_, err := Redis.Ping(ctx).Result()
	if err != nil {
		if err == context.DeadlineExceeded || err == context.Canceled {
			panic("Ping redis time out")
		}
		panic("Error ping redis")
	}
}

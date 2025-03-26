package ioc

import (
	"github.com/LEILEI0628/GoWeb/config"
	"github.com/redis/go-redis/v9"
)

func InitRedis() redis.Cmdable {
	return redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Addr,
		Password: "",
		DB:       0,
	})
}

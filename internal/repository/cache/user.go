package cache

import (
	"fmt"
	"github.com/LEILEI0628/GinPro/middleware/cache"
	"github.com/LEILEI0628/GoWeb/internal/domain"
	"github.com/redis/go-redis/v9"
	"time"
)

type UserCache = cachex.Cache[domain.User, int64]

func InitUserCache(redisClient *redis.Client, expiration time.Duration) *UserCache {
	// 用户缓存初始化方法
	userKeyFunc := func(id int64) string {
		return fmt.Sprintf("user:info:%d", id)
	}
	return cachex.NewCache[domain.User, int64](redisClient, expiration, userKeyFunc)

}

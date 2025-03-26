package cache

import (
	"fmt"
	cachex "github.com/LEILEI0628/GinPro/middleware/cache"
	"github.com/LEILEI0628/GoWeb/internal/domain"
	"github.com/redis/go-redis/v9"
	"time"
)

type UserCache = cachex.RedisCache[domain.User, int64]

func NewUserCache(client redis.Cmdable) *UserCache {
	expiration := time.Minute * 15
	// 用户缓存初始化方法
	userKeyFunc := func(id int64) string {
		return fmt.Sprintf("user:info:%d", id)
	}
	return cachex.NewRedisCache[domain.User, int64](client, expiration, userKeyFunc)

}

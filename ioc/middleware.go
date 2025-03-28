package ioc

import (
	"github.com/LEILEI0628/GinPro/middleware/limiter"
	"github.com/LEILEI0628/GoWeb/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"time"
)

func InitMiddleware(redisClient redis.Cmdable) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.ResolveCORS(), // 解决跨域问题
		middleware.JWT(),
		limiter.NewBuilder(limiter.NewRedisSlidingWindowLimiter(redisClient, time.Second, 1000)).
			Build(), // 限流（滑动窗口算法）1000/1s 使用redis统计请求数量
	}
}

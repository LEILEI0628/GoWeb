package ioc

import (
	"github.com/LEILEI0628/GinPro/middleware/limiter"
	"github.com/LEILEI0628/GinPro/middleware/ratelimit"
	"github.com/LEILEI0628/GoWeb/internal/middleware"
	"github.com/LEILEI0628/GoWeb/internal/web"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"time"
)

func InitGin(middleware []gin.HandlerFunc, routers *web.Routers) *gin.Engine {
	server := gin.Default()
	server.Use(middleware...)
	routers.RegisterRouters(server)
	return server
}

func InitMiddleware(redisClient redis.Cmdable) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.ResolveCORS(), // 解决跨域问题
		middleware.JWT(),
		ratelimit.NewBuilder(limiter.NewRedisSlidingWindowLimiter(redisClient, time.Second, 1000)).
			Build(), // 限流（滑动窗口算法）使用redis统计请求数量
	}
}

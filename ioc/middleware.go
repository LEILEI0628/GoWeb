package ioc

import (
	"github.com/LEILEI0628/GinPro/middleware/limiter"
	"github.com/LEILEI0628/GoWeb/internal/middleware"
	"github.com/gin-gonic/gin"
)

func InitMiddleware(limiterParam limiter.Limiter) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		middleware.ResolveCORS(), // 解决跨域问题
		middleware.JWT(),
		middleware.Limiter(limiterParam),
	}
}

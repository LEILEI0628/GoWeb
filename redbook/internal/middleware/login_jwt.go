package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang-web-learn/redbook/internal/web/user"
	"net/http"
	"strings"
	"time"
)

type LoginMiddlewareJWTBuilder struct { // 使用Build模式时不要对顺序进行任何的设定
	paths []string
}

func NewLoginMiddlewareJWTBuilder() *LoginMiddlewareJWTBuilder {
	return &LoginMiddlewareJWTBuilder{}
}

func (loginJWTBuilder LoginMiddlewareJWTBuilder) Build() gin.HandlerFunc {
	return func(context *gin.Context) {
		for _, path := range loginJWTBuilder.paths {
			if context.Request.URL.Path == path {
				return // 无需登录校验
			}
		}

		// 使用JWT校验
		tokenHeader := context.GetHeader("Authorization")
		if tokenHeader == "" {
			// 还未登录
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		segs := strings.Split(tokenHeader, " ") // Authorization规范token前会有一个“Bearer ”
		if len(segs) != 2 {
			// token格式错误
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		tokenStr := segs[1]
		claims := &web.UserClaims{}
		// ParseWithClaims方法的claims参数一定要传指针，方法会对claims进行修改
		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("7x9FpL2QaZ8rT4wY6vBcN1mK3jH5gD7s"), nil
		})
		if err != nil || !token.Valid || token == nil || claims.UID == 0 { // 过期Valid为false
			// token检验错误
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if claims.UserAgent != context.Request.UserAgent() { // 可以只匹配部分内容，减少误操作的可能
			// 严重的安全问题
			// TODO 监控
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// 通过校验
		// token续约（每10分钟）
		now := time.Now()
		if claims.ExpiresAt.Sub(now) < time.Minute*10 {
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(12 * time.Hour))              // 12小时后过期
			newTokenStr, err := token.SignedString([]byte("7x9FpL2QaZ8rT4wY6vBcN1mK3jH5gD7s")) // 重新生成token
			if err != nil {
				// 无需中断程序运行
				// TODO 记录日志：续约失败
			}
			context.Header("x-jwt-token", newTokenStr)
		}
		context.Set("claims", claims)

	}
}

func (loginJWTBuilder LoginMiddlewareJWTBuilder) IgnorePaths(path string) *LoginMiddlewareJWTBuilder {
	loginJWTBuilder.paths = append(loginJWTBuilder.paths, path)
	return &loginJWTBuilder
}

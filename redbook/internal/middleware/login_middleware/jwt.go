package login_middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	web "golang-web-learn/redbook/internal/web/user"
	"net/http"
	"strings"
	"time"
)

// BuildByJWT 终结方法（使用JWT进行校验）verificationKey:校验秘钥；expiresTime：JWT过期时间；leftTime：续约剩余时间
func (loginMiddlewareBuilder *LoginMiddlewareBuilder) BuildByJWT(verificationKey string, expiresTime time.Duration, leftTime time.Duration) gin.HandlerFunc {
	return func(context *gin.Context) {
		for _, path := range loginMiddlewareBuilder.ignorePaths {
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
			return []byte(verificationKey), nil
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
		// token续约（还剩leftTime时）
		now := time.Now()
		if claims.ExpiresAt.Sub(now) < leftTime {
			claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(expiresTime)) // expiresTime后过期
			newTokenStr, err := token.SignedString([]byte(verificationKey))    // 重新生成token
			if err != nil {
				// 无需中断程序运行
				// TODO 记录日志：续约失败
			}
			context.Header("x-login_middleware-token", newTokenStr)
		}
		context.Set("claims", claims)
	}
}

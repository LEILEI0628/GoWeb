package login_middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// BuildBySession 终结方法（使用Session进行校验）
func (loginMiddlewareBuilder *LoginMiddlewareBuilder) BuildBySession(maxAgeSec int, leftTime time.Duration) gin.HandlerFunc {
	return func(context *gin.Context) {
		for _, path := range loginMiddlewareBuilder.ignorePaths {
			if context.Request.URL.Path == path {
				return // 无需登录校验
			}
		}

		session := sessions.Default(context)

		id := session.Get("userId")
		if id == nil {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// TODO 封装插件
		// 实现定时刷新token操作
		now := time.Now().UnixMilli()
		updateTime := session.Get("updateTime")
		session.Set("userId", id)
		err := session.Save()
		if err != nil {
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if updateTime == nil {
			// 还没刷新过
			session.Set("updateTime", now)
			session.Options(sessions.Options{
				MaxAge: maxAgeSec,
			})
			err := session.Save()
			if err != nil {
				context.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			return
		}
		updateTimeVal, ok := updateTime.(int64)
		if !ok {
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if now-updateTimeVal > leftTime.Milliseconds() { // 一分钟刷新一次
			session.Set("updateTime", now)
			session.Options(sessions.Options{
				MaxAge: maxAgeSec,
			})
			err := session.Save()
			if err != nil {
				context.AbortWithStatus(http.StatusInternalServerError)
				return
			}
			return
		}

	}
}

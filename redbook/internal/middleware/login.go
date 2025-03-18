package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginMiddlewareBuilder struct { // 使用Build模式时不要对顺序进行任何的设定
	paths []string
}

func NewLoginMiddlewareBuilder() *LoginMiddlewareBuilder {
	return &LoginMiddlewareBuilder{}
}

func (loginMiddlewareBuilder LoginMiddlewareBuilder) Build() gin.HandlerFunc {
	return func(context *gin.Context) {
		for _, path := range loginMiddlewareBuilder.paths {
			if context.Request.URL.Path == path {
				return // 无需登录校验
			}
		}

		session := sessions.Default(context)
		if session.Get("userId") == nil {
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}

func (loginMiddlewareBuilder LoginMiddlewareBuilder) IgnorePaths(path string) *LoginMiddlewareBuilder {
	loginMiddlewareBuilder.paths = append(loginMiddlewareBuilder.paths, path)
	return &loginMiddlewareBuilder
}

func (loginMiddlewareBuilder LoginMiddlewareBuilder) IncludePaths(path string) *LoginMiddlewareBuilder {
	loginMiddlewareBuilder.paths = append(loginMiddlewareBuilder.paths, path)
	return &loginMiddlewareBuilder
}

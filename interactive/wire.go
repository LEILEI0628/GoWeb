//go:build wireinject

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitWebServer() *gin.Engine {
	wire.Build(
		// 初始化最基础的第三方依赖
		interactiveSvcProvider,
	)
	return new(gin.Engine)
}

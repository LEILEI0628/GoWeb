//go:build wireinject

package main

import (
	"github.com/LEILEI0628/GoWeb/internal/repository"
	"github.com/LEILEI0628/GoWeb/internal/repository/cache"
	"github.com/LEILEI0628/GoWeb/internal/repository/dao"
	"github.com/LEILEI0628/GoWeb/internal/service"
	"github.com/LEILEI0628/GoWeb/internal/web/handler"
	"github.com/LEILEI0628/GoWeb/internal/web/router"
	"github.com/LEILEI0628/GoWeb/ioc"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitWebServer() *gin.Engine {
	wire.Build(
		// 最基础的第三方依赖
		ioc.InitDB, ioc.InitRedis,
		// 初始化DAO
		dao.NewUserDAO,

		cache.NewUserCache,

		repository.NewUserRepository,

		service.NewUserService,

		handler.NewUserHandler,
		router.NewUserRouters,
		ioc.InitMiddleware,
		ioc.InitGin,
	)
	return new(gin.Engine)
}

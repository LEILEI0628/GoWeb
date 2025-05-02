//go:build wireinject

package main

import (
	"github.com/LEILEI0628/GoWeb/internal/repository"
	article2 "github.com/LEILEI0628/GoWeb/internal/repository/article"
	"github.com/LEILEI0628/GoWeb/internal/repository/cache"
	"github.com/LEILEI0628/GoWeb/internal/repository/dao"
	"github.com/LEILEI0628/GoWeb/internal/repository/dao/article"
	"github.com/LEILEI0628/GoWeb/internal/service"
	"github.com/LEILEI0628/GoWeb/internal/web"
	"github.com/LEILEI0628/GoWeb/internal/web/handler"
	"github.com/LEILEI0628/GoWeb/internal/web/router"
	"github.com/LEILEI0628/GoWeb/ioc"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

func InitWebServer() *gin.Engine {
	wire.Build(
		// 初始化最基础的第三方依赖
		ioc.InitDB, ioc.InitRedis,

		// 初始化DAO
		dao.NewUserDAO,
		article.NewArticleDAO,
		// 初始化Cache
		cache.NewUserCache,
		// 初始化Repository
		repository.NewUserRepository,
		article2.NewArticleRepository,
		// 初始化Service
		ioc.InitUserService,
		service.NewArticleService,
		// 初始化Handler
		handler.NewUserHandler,
		handler.NewArticleHandler,
		// 初始化Routers
		router.NewUserRouters,
		router.NewArticleRouters,

		// 初始化Routers、中间件、server
		web.NewRegisterRouters,
		ioc.InitLimiter,
		ioc.InitGlobalLogger,
		ioc.InitMiddleware,
		ioc.InitGin,
	)
	return new(gin.Engine)
}

package main

import (
	"github.com/LEILEI0628/GinPro/middleware/limiter"
	"github.com/LEILEI0628/GinPro/middleware/ratelimit"
	"github.com/LEILEI0628/GoWeb/config"
	"github.com/LEILEI0628/GoWeb/internal/web"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"time"
)

// 编译命令：GOOS=linux GOARCH=arm go build -o goweb .
func main() {
	server := gin.Default()

	redisClient := initRedis()
	db := initDB()

	// 限流（滑动窗口算法）使用redis统计请求数量
	server.Use(ratelimit.NewBuilder(limiter.NewRedisSlidingWindowLimiter(redisClient, time.Second, 1000)).
		Build())

	web.NewInitWeb(server, db, redisClient).RegisterRouters()
	// 下列代码已被封装
	// START
	//userHandler := web.UserHandler{}
	//server.GET("/users/profile", userHandler.Profile)
	//server.POST("/users/edit", userHandler.Edit)
	// REST风格
	//server.GET("/users/profile/:id", userHandler.Profile)
	//server.POST("/users/edit/:id", userHandler.Edit)
	// END
	server.GET("/hello", func(context *gin.Context) {
		context.String(http.StatusOK, "hello world")
	})
	server.Run(":8080")
}

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     config.Config.Redis.Addr,
		Password: "",
		DB:       0,
	})
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.Config.DB.DSN))
	if err != nil {
		// panic相当于整个goroutine结束
		// panic只会出现在初始化的过程中（一旦初始化出错，就没必要启动了）
		panic(err)
	}

	// 建表
	//err = dao.InitUserTable(db)
	//if err != nil {
	//	panic(err)
	//}

	return db
}

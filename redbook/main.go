package main

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"golang-web-learn/redbook/internal/web"
	"golang-web-learn/redbook/pkg/ginx/middleware/ratelimit"
	"golang-web-learn/redbook/pkg/limiter"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func main() {
	server := gin.Default()

	db := initDB()

	redisClient := initRedis()

	// 限流（滑动窗口算法）使用redis统计请求数量
	server.Use(ratelimit.
		NewBuilder(limiter.NewRedisSlidingWindowLimiter(redisClient, time.Second, 1000)).
		Build())

	web.RegisterRouters(server, db)
	// 下列代码已被封装
	// START
	//userHandler := web.UserHandler{}
	//server.GET("/users/profile", userHandler.Profile)
	//server.POST("/users/edit", userHandler.Edit)
	// REST风格
	//server.GET("/users/profile/:id", userHandler.Profile)
	//server.POST("/users/edit/:id", userHandler.Edit)
	// END
	server.Run(":8080")
}

func initRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func initDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:20010628@tcp(localhost:13316)/redbook"))
	if err != nil {
		// panic相当于整个goroutine结束
		// panic只会出现在初始化的过程中（一旦初始化出错，就没必要启动了）
		panic(err)
	}

	// 建表
	//err := dao.InitUserTable(db)
	//if err != nil {
	//	panic(err)
	//}

	return db
}

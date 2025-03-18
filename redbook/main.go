package main

import (
	"github.com/gin-gonic/gin"
	"golang-web-learn/redbook/internal/web"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	server := gin.Default()

	db := initDB()

	web.RegisterRouters(server, db)
	// 下列代码已被封装
	// START
	//userHandler := web.UserHandler{}
	//server.POST("/users/signIn", userHandler.SignIn)
	//server.POST("/users/signUp", userHandler.SignUp)
	//server.GET("/users/profile", userHandler.Profile)
	//server.POST("/users/edit", userHandler.Edit)
	// REST风格
	//server.GET("/users/profile/:id", userHandler.Profile)
	//server.POST("/users/edit/:id", userHandler.Edit)
	// END
	server.Run(":8080")
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

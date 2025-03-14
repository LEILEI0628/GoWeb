package main

import (
	"github.com/gin-gonic/gin"
	"golang-web-learn/redbook/internal/web"
)

func main() {
	server := gin.Default()
	web.RegisterRouters(server)
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

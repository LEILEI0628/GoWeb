package main

// 编译命令：GOOS=linux GOARCH=arm go build -o goweb .
func main() {
	//server := gin.Default()
	//redisClient := initRedis()
	//db := initDB()
	server := InitWebServer()

	//web.NewInitWeb(server).RegisterRouters()
	// 下列代码已被封装
	// START
	//userHandler := web.UserHandler{}
	//server.GET("/users/profile", userHandler.Profile)
	//server.POST("/users/edit", userHandler.Edit)
	// REST风格
	//server.GET("/users/profile/:id", userHandler.Profile)
	//server.POST("/users/edit/:id", userHandler.Edit)
	// END
	//server.GET("/hello", func(context *gin.Context) {
	//	context.String(http.StatusOK, "hello world")
	//})
	server.Run(":8080")
}

package main

// 编译命令：GOOS=linux GOARCH=arm go build -o goweb .
func main() {
	server := InitWebServer()
	//server.GET("/hello", func(context *gin.Context) {
	//	context.String(http.StatusOK, "hello world")
	//})
	server.Run(":8080")
}

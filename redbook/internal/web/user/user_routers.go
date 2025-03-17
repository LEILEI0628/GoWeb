package user

import (
	"github.com/gin-gonic/gin"
)

var userHandler *UserHandler

// Routers User相关的路由
func Routers(server *gin.Engine) {
	userHandler = NewUserHandler()
	// 分组路由
	userGroup := server.Group("/users")
	signUpRouter(userGroup)
	signInRouter(userGroup)
	profileRouter(userGroup)
	editRouter(userGroup)

}

func signUpRouter(userGroup *gin.RouterGroup) {
	userGroup.POST("/signup", userHandler.SignUp)
}

func signInRouter(userGroup *gin.RouterGroup) {
	userGroup.POST("/signin", userHandler.SignIn)
}

func editRouter(userGroup *gin.RouterGroup) {
	userGroup.POST("/edit", userHandler.Edit)

}

func profileRouter(userGroup *gin.RouterGroup) {
	userGroup.GET("/profile", userHandler.Profile)

}

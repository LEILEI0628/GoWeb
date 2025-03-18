package user

import (
	"github.com/gin-gonic/gin"
)

// UserRouters User相关的路由
type UserRouters struct {
	userHandler *UserHandler
	server      *gin.Engine
}

func NewUserRouters(userHandler *UserHandler, server *gin.Engine) *UserRouters {
	return &UserRouters{userHandler: userHandler, server: server}
}

func (userRouters UserRouters) RegisterUserRouters() {
	// 分组路由
	userGroup := userRouters.server.Group("/users")
	userRouters.signUpRouter(userGroup)
	userRouters.signInRouter(userGroup)
	userRouters.profileRouter(userGroup)
	userRouters.editRouter(userGroup)

}

func (userRouters UserRouters) signUpRouter(userGroup *gin.RouterGroup) {
	userGroup.POST("/signup", userRouters.userHandler.SignUp)
}

func (userRouters UserRouters) signInRouter(userGroup *gin.RouterGroup) {
	userGroup.POST("/signin", userRouters.userHandler.SignIn)
}

func (userRouters UserRouters) editRouter(userGroup *gin.RouterGroup) {
	userGroup.POST("/edit", userRouters.userHandler.Edit)

}

func (userRouters UserRouters) profileRouter(userGroup *gin.RouterGroup) {
	userGroup.GET("/profile", userRouters.userHandler.Profile)

}

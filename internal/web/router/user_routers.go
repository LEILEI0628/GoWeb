package router

import (
	"github.com/LEILEI0628/GoWeb/internal/web/handler"
	"github.com/gin-gonic/gin"
)

// UserRouters User相关的路由
type UserRouters struct {
	userHandler *handler.UserHandler
}

func NewUserRouters(userHandler *handler.UserHandler) *UserRouters {
	return &UserRouters{userHandler: userHandler}
}

func (userRouters UserRouters) RegisterUserRouters(server *gin.Engine) {
	// 分组路由
	userGroup := server.Group("/users")
	userRouters.signUpRouter(userGroup)
	userRouters.signInRouter(userGroup)
	userRouters.signOutRouter(userGroup)
	userRouters.profileRouter(userGroup)
	userRouters.editRouter(userGroup)

}

func (userRouters UserRouters) signUpRouter(userGroup *gin.RouterGroup) {
	userGroup.POST("/signup", userRouters.userHandler.SignUp)
}

func (userRouters UserRouters) signInRouter(userGroup *gin.RouterGroup) {
	userGroup.POST("/login", userRouters.userHandler.SignInByJWT)
}

func (userRouters UserRouters) signOutRouter(userGroup *gin.RouterGroup) {
	userGroup.GET("/logout", userRouters.userHandler.SignOut)

}

func (userRouters UserRouters) editRouter(userGroup *gin.RouterGroup) {
	userGroup.POST("/edit", userRouters.userHandler.Edit)

}

func (userRouters UserRouters) profileRouter(userGroup *gin.RouterGroup) {
	userGroup.GET("/profile", userRouters.userHandler.ProfileByJWT)

}

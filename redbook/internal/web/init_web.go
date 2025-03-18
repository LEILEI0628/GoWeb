package web

// 另一种做法是将该文件放置在main.go同级
import (
	"github.com/gin-gonic/gin"
	"golang-web-learn/redbook/internal/middleware"
	"golang-web-learn/redbook/internal/repository"
	"golang-web-learn/redbook/internal/repository/dao"
	"golang-web-learn/redbook/internal/service"
	"golang-web-learn/redbook/internal/web/user"
	"gorm.io/gorm"
)

// RegisterRouters 注册路由
func RegisterRouters(server *gin.Engine, db *gorm.DB) {
	//server.Use(func(c *gin.Context) { // Use作用于全部路由
	//	fmt.Println("自定义的middleware")
	//})
	middleware.ResolveCROS(server) // 解决跨域问题
	initUserRouters(server, db).RegisterUserRouters()
}

func initUserRouters(server *gin.Engine, db *gorm.DB) *user.UserRouters {
	userDAO := dao.NewUserDAO(db)
	userRepository := repository.NewUserRepository(userDAO)
	userService := service.NewUserService(userRepository)
	userHandler := user.NewUserHandler(userService)
	return user.NewUserRouters(userHandler, server)
}

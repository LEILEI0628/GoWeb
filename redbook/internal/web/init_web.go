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

	globalMiddleware := middleware.NewGlobalMiddlewareBuilder()
	server.Use(globalMiddleware.ResolveCORS()) // 解决跨域问题
	//server.Use(globalMiddleware.Session())     // 添加session（cookie中）（存储方式在方法中自定义）
	//server.Use(middleware.NewLoginMiddlewareBuilder(). // 校验session
	server.Use(middleware.NewLoginMiddlewareJWTBuilder(). // 校验JWT
								IgnorePaths("/users/login"). // 链式调用，不同的server可定制（扩展性）
								IgnorePaths("/users/signup").
								Build()) // Builder模式为了解决复杂结构构建问题

	initUserRouters(server, db).RegisterUserRouters()

}

func initUserRouters(server *gin.Engine, db *gorm.DB) *web.UserRouters {
	userDAO := dao.NewUserDAO(db)
	userRepository := repository.NewUserRepository(userDAO)
	userService := service.NewUserService(userRepository)
	userHandler := web.NewUserHandler(userService)
	return web.NewUserRouters(userHandler, server)
}

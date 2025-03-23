package web

// 另一种做法是将该文件放置在main.go同级
import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"golang-web-learn/redbook/internal/middleware"
	"golang-web-learn/redbook/internal/repository"
	"golang-web-learn/redbook/internal/repository/cache"
	"golang-web-learn/redbook/internal/repository/dao"
	"golang-web-learn/redbook/internal/service"
	"golang-web-learn/redbook/internal/web/user"
	"gorm.io/gorm"
	"time"
)

type InitWeb struct {
	server      *gin.Engine
	db          *gorm.DB
	redisClient *redis.Client
}

func NewInitWeb(server *gin.Engine, db *gorm.DB, redisClient *redis.Client) *InitWeb {
	return &InitWeb{server: server, db: db, redisClient: redisClient}
}

// RegisterRouters 注册路由
func (initWeb InitWeb) RegisterRouters() {
	//server.Use(func(c *gin.Context) { // Use作用于全部路由
	//	fmt.Println("自定义的middleware")
	//})

	globalMiddleware := middleware.NewGlobalMiddlewareBuilder()
	initWeb.server.Use(globalMiddleware.ResolveCORS()) // 解决跨域问题
	//server.Use(globalMiddleware.Session())     // 添加session（cookie中）（存储方式在方法中自定义）
	//server.Use(middleware.NewLoginMiddlewareBuilder(). // 校验session
	initWeb.server.Use(middleware.NewLoginMiddlewareJWTBuilder(). // 校验JWT
									IgnorePaths("/users/login"). // 链式调用，不同的server可定制（扩展性）
									IgnorePaths("/users/signup").
									IgnorePaths("/hello").
									Build()) // Builder模式为了解决复杂结构构建问题

	initWeb.initUserRouters().RegisterUserRouters()

}

func (initWeb InitWeb) initUserRouters() *web.UserRouters {
	userDAO := dao.NewUserDAO(initWeb.db)
	userCache := cache.NewUserCache(initWeb.redisClient, time.Minute*15)
	userRepository := repository.NewUserRepository(userDAO, userCache)
	userService := service.NewUserService(userRepository)
	userHandler := web.NewUserHandler(userService)
	return web.NewUserRouters(userHandler, initWeb.server)
}

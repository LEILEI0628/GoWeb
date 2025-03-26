package web

// 另一种做法是将该文件放置在main.go同级
import (
	"github.com/gin-gonic/gin"
)

type InitWeb struct {
	server *gin.Engine
	//db          *gorm.DB
	//redisClient *redis.Client
}

func NewInitWeb(server *gin.Engine) *InitWeb {
	//return &InitWeb{server: server, db: db, redisClient: redisClient}
	return &InitWeb{server: server}
}

// RegisterRouters 注册路由
func (initWeb InitWeb) RegisterRouters() {
	//server.Use(func(c *gin.Context) { // Use作用于全部路由
	//	fmt.Println("自定义的middleware")
	//})

	//initWeb.server.Use(middleware.ResolveCORS()) // 解决跨域问题
	//
	//initWeb.server.Use(jwtx.NewBuilder( // 校验JWT
	//	jwtx.WithVerificationKey("7x9FpL2QaZ8rT4wY6vBcN1mK3jH5gD7s"),
	//	jwtx.WithExpiresTime(time.Hour*12),
	//	jwtx.WithLeftTime(time.Minute*10)).
	//	IgnorePaths("/users/login"). // 链式调用，不同的server可定制（扩展性）
	//	IgnorePaths("/users/signup").
	//	IgnorePaths("/hello").Build()) // Builder模式为了解决复杂结构构建问题

	//sessionConfig := session.Config{
	//	StorageType: session.Redis,
	//	AuthKey:     []byte("7x9FpL2QaZ8rT4wY6vBcN1mK3jH5gD7s"),
	//	EncryptKey:  []byte("qW3eRtY8uI0oP9aSsDfGhJkL4zXcV6bN"),
	//	RedisOpts: session.RedisOpts{
	//		MaxIdle:  16,
	//		Network:  "tcp",
	//		Addr:     config.Config.Redis.Addr,
	//		Password: "",
	//	},
	//}
	//initWeb.server.Use(session.SessionStore(sessionConfig)) // 添加session（cookie中）（存储使用redis）
	//initWeb.server.Use(session.NewBuilder().                // 校验session
	//Build(60*60, time.Minute)) // 使用session校验

	//initWeb.initUserRouters().RegisterUserRouters()

}

//func (initWeb InitWeb) initUserRouters() *web.UserRouters {
//	// 强耦合初始化
//	userDAO := dao.NewUserDAO(initWeb.db)
//	userCache := cache.InitUserCache(initWeb.redisClient, time.Minute*15)
//	userRepository := repository.NewUserRepository(userDAO, userCache)
//	userService := service.NewUserService(userRepository)
//	userHandler := web.NewUserHandler(userService)
//	return web.NewUserRouters(userHandler, initWeb.server)
//}

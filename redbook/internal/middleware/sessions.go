package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func (b GlobalMiddlewareBuilder) Session() gin.HandlerFunc {
	// cookie和memstore的NewStore()第一个参数为authentication key，第二个参数为encryption key，推荐32or64位
	//store := cookie.NewStore([]byte("7x9FpL2QaZ8rT4wY6vBcN1mK3jH5gD7s"), []byte("qW3eRtY8uI0oP9aSsDfGhJkL4zXcV6bN")) // 基于cookie的实现
	//store := memstore.NewStore([]byte("7x9FpL2QaZ8rT4wY6vBcN1mK3jH5gD7s"), []byte("qW3eRtY8uI0oP9aSsDfGhJkL4zXcV6bN")) // 基于内存的实现

	// 基于redis实现
	// 第一个参数是最大空闲连接数量，第二个参数是连接方式，第三四个参数是连接信息和密码，第五六个是key
	store, err := redis.NewStore(16, "tcp", "localhost:6379", "",
		[]byte("7x9FpL2QaZ8rT4wY6vBcN1mK3jH5gD7s"), []byte("7x9FpL2QaZ8rT4wY6vBcN1mK3jH5gD7s"))
	if err != nil {
		panic(err)
	}
	return sessions.Sessions("rb_session", store)

}

package web

import (
	"github.com/LEILEI0628/GoWeb/internal/web/router"
	"github.com/gin-gonic/gin"
)

type Routers struct {
	userRouters *router.UserRouters
}

func NewRouters(userRouters *router.UserRouters) *Routers { // 后期需要不断添加routers
	return &Routers{userRouters: userRouters}
}

func (routers *Routers) RegisterRouters(server *gin.Engine) {
	routers.userRouters.RegisterUserRouters(server) // 后期需要不断添加routers
}

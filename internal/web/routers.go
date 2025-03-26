package web

import (
	"github.com/LEILEI0628/GoWeb/internal/web/router"
	"github.com/gin-gonic/gin"
)

type Routers interface {
	RegisterRouters(server *gin.Engine)
}

type RegisterRouters struct {
	routers []Routers
}

func NewRegisterRouters(userRouters *router.UserRouters) *RegisterRouters { // 后期需要不断添加routers
	routers := []Routers{userRouters}
	return &RegisterRouters{routers: routers}
}

func (rr *RegisterRouters) Register(server *gin.Engine) {
	for _, v := range rr.routers {
		v.RegisterRouters(server)
	}
}

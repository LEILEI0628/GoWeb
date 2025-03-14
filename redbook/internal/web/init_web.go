package web

// 另一种做法是将该文件放置在main.go同级
import (
	"github.com/gin-gonic/gin"
	"golang-web-learn/redbook/internal/web/user"
)

// RegisterRouters 注册路由
func RegisterRouters(server *gin.Engine) {
	user.Routers(server)
}

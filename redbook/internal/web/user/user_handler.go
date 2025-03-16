package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Handler User相关的业务处理
type Handler struct{}

func (userHandler *Handler) SignUp(context *gin.Context) {
	// 内部结构体：确保只有本方法能访问
	type SignUpReq struct {
		Email           string `json:"email"` // `json:"email"`：标签
		Password        string `json:"password"`
		ConfirmPassword string `json:"confirmPassword"`
	}

	var req SignUpReq
	// Bind方法会根据Content-Type来解析数据到json中，出错直接返回400状态码
	if err := context.Bind(&req); err != nil { // 注意：&req
		return
	}

	// 校验操作
	const ( // 就近原则和最小化作用域原则
		emailRegexPattern    = ""
		passwordRegexPattern = ""
	)

	//ok, err := regexp.Match(emailRegexPattern, []byte(req.Email))
	//if err != nil {
	//	// 不要把后端错误信息传到前端！err.Error()
	//	// TODO 记录日志
	//	context.String(http.StatusOK, "系统错误") // 可以直观的看到请求是否到达服务器
	//	//context.String(http.StatusInternalServerError, "系统错误") // 更符合http规范
	//	return
	//}
	//if !ok {
	//	context.String(http.StatusOK, "邮箱格式错误")
	//	//context.String(http.StatusBadRequest, "邮箱格式错误")
	//	return
	//}
	// TODO 其他校验操作
	// TODO ORM操作
	fmt.Printf("%v\n", req)
	context.String(http.StatusOK, "注册成功！")
}

func (userHandler *Handler) SignIn(context *gin.Context) {

}

func (userHandler *Handler) Edit(context *gin.Context) {

}

func (userHandler *Handler) Profile(context *gin.Context) {

}

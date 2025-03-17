package user

import (
	"fmt"
	regexp "github.com/dlclark/regexp2" // 自带的regexp无法处理复杂正则
	"github.com/gin-gonic/gin"
	"golang-web-learn/redbook/internal/service"
	"net/http"
)

// Handler User相关的业务处理
type UserHandler struct {
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
	userService *service.UserService
}

func NewUserHandler() *UserHandler {
	const ( // 就近原则和最小化作用域原则
		emailRegexPattern = `^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$` // ``比""简洁（无需转义）
		//emailRegexPattern = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
		passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[@$!%*#?&.])[A-Za-z\d@$!%*#?&.]{8,}$`
	)
	emailExp := regexp.MustCompile(emailRegexPattern, regexp.None)
	passwordExp := regexp.MustCompile(passwordRegexPattern, regexp.None)

	return &UserHandler{
		emailExp:    emailExp,
		passwordExp: passwordExp,
	}
}

func (userHandler *UserHandler) SignUp(context *gin.Context) {
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
	ok, err := userHandler.emailExp.MatchString(req.Email)
	if err != nil {
		// 此处的err是正则表达式错误（使用regexp2时只有超时才会出现err）
		// 不要把后端错误信息（err.Error()）传到前端！
		// TODO 记录日志
		context.String(http.StatusOK, "系统错误") // 可以直观的看到请求是否到达服务器
		//context.String(http.StatusInternalServerError, "系统错误") // RESTful风格更符合http规范
		return
	}
	if !ok {
		context.String(http.StatusOK, "邮箱格式错误")
		//context.String(http.StatusBadRequest, "邮箱格式错误")
		return
	}

	if req.Password != req.ConfirmPassword {
		// TODO 记录日志
		context.String(http.StatusOK, "两次输入的密码不一致")
	}

	ok, err = userHandler.passwordExp.MatchString(req.Password)
	if err != nil {
		// TODO 记录日志
		context.String(http.StatusOK, "系统错误")
	}

	if !ok {
		context.String(http.StatusOK, "密码必须大于8位且包含数字、特殊字符")
		return
	}

	// TODO ORM操作
	fmt.Printf("%v\n", req)
	context.String(http.StatusOK, "注册成功！")
}

func (userHandler *UserHandler) SignIn(context *gin.Context) {

}

func (userHandler *UserHandler) Edit(context *gin.Context) {

}

func (userHandler *UserHandler) Profile(context *gin.Context) {

}

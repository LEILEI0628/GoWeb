package user

import (
	"errors"
	regexp "github.com/dlclark/regexp2" // 自带的regexp无法处理复杂正则
	"github.com/gin-gonic/gin"
	"golang-web-learn/redbook/internal/domain"
	"golang-web-learn/redbook/internal/service"
	"net/http"
)

// Handler User相关的业务处理
type UserHandler struct {
	emailExp    *regexp.Regexp
	passwordExp *regexp.Regexp
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler { // 使用此方法可以提示忘记传参
	const ( // 就近原则和最小化作用域原则
		emailRegexPattern = `^\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*$` // ``比""简洁（无需转义）
		//emailRegexPattern = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
		passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[@$!%*#?&.])[A-Za-z\d@$!%*#?&.]{8,72}$`
	)
	emailExp := regexp.MustCompile(emailRegexPattern, regexp.None)
	passwordExp := regexp.MustCompile(passwordRegexPattern, regexp.None)

	return &UserHandler{
		emailExp:    emailExp,
		passwordExp: passwordExp,
		userService: userService,
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
	err := userHandler.checkMessage(context, req.Email, req.Password, req.ConfirmPassword)
	if err != nil {
		return
	}

	err = userHandler.userService.SignUp(context.Request.Context(), domain.User{Email: req.Email, Password: req.Password})
	if errors.Is(err, service.ErrUserEmailDuplicated) {
		// 使用service.Err...而不直接使用dao.Err...是为了防止跨层耦合
		context.String(http.StatusOK, "邮箱冲突")
		return
	}

	if err != nil {
		// TODO 记录日志
		context.String(http.StatusOK, "系统错误")
		return
	}

	context.String(http.StatusOK, "注册成功！")
}

func (userHandler *UserHandler) SignIn(context *gin.Context) {

}

func (userHandler *UserHandler) Edit(context *gin.Context) {

}

func (userHandler *UserHandler) Profile(context *gin.Context) {

}

func (userHandler *UserHandler) checkMessage(context *gin.Context, email string, password string, confirmPassword string) error {
	ok, err := userHandler.emailExp.MatchString(email)
	if err != nil {
		// 此处的err是正则表达式错误（使用regexp2时只有超时才会出现err）
		// 不要把后端错误信息（err.Error()）传到前端！
		// TODO 记录日志
		context.String(http.StatusOK, "系统错误") // 可以直观的看到请求是否到达服务器
		//context.String(http.StatusInternalServerError, "系统错误") // RESTful风格更符合http规范
		return err
	}
	if !ok {
		context.String(http.StatusOK, "邮箱格式错误")
		//context.String(http.StatusBadRequest, "邮箱格式错误")
		return errors.New("邮箱格式错误")
	}

	if password != confirmPassword {
		// TODO 记录日志
		context.String(http.StatusOK, "两次输入的密码不一致")
		return errors.New("两次输入的密码不一致")
	}

	ok, err = userHandler.passwordExp.MatchString(password)
	if err != nil {
		// TODO 记录日志
		context.String(http.StatusOK, "系统错误")
		return err
	}

	if !ok {
		context.String(http.StatusOK, "密码必须大于8位且包含数字、特殊字符")
		return errors.New("密码格式错误")
	}
	return nil
}

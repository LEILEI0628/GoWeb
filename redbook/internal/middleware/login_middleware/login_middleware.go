package login_middleware

type LoginMiddlewareBuilder struct { // 使用Build模式时不要对顺序进行任何的设定
	ignorePaths []string
}

func NewLoginMiddlewareBuilder() *LoginMiddlewareBuilder {
	return &LoginMiddlewareBuilder{}
}

// Build 终结方法，请在最后调用
//func (loginMiddlewareBuilder *LoginMiddlewareBuilder) Build() gin.HandlerFunc {
//	if loginMiddlewareBuilder.calibrationType == 1 { // JWT校验
//	}
//	if loginMiddlewareBuilder.calibrationType == 2 { // session校验
//	}
//}

//func (loginMiddlewareBuilder *LoginMiddlewareBuilder) checkWithJWT() gin.HandlerFunc {
//
//}

// IgnorePaths 要忽略的路径
func (loginMiddlewareBuilder *LoginMiddlewareBuilder) IgnorePaths(path string) *LoginMiddlewareBuilder {
	// 中间方法
	// 注：方法接收器使用值接收器时每次调用方法都会创建一个副本，当进行取地址操作时可以实现功能，
	// 返回的是新副本的指针，但原实例并未更改，这也就造成了资源浪费，因此强烈建议使用指针接收器
	loginMiddlewareBuilder.ignorePaths = append(loginMiddlewareBuilder.ignorePaths, path)
	return loginMiddlewareBuilder
}

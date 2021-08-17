package request

type LoginInput struct {
	Account string `json:"account" v:"required#登录账号不能为空"`
	Password string `json:"password"  v:"required#登录密码不能为空"`
}

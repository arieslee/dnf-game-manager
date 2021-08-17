package response

const (
	ErrorServerError        = 50000 // 服务器内部错误
	ErrorServerUnrealized   = 50001 // 网络未实现
	ErrorServerBusy         = 50002 // 系统繁忙
	ErrorServiceUnavailable = 50003 // 服务不可用
	ErrorGatewayTimeout     = 50004 // 网关超时
	ErrorUserForbidden      = 50005 // 会员状态为禁止

	ErrorBadRequest        = 40000 // 非法请求
	ErrorAccessTokenError  = 40001 // access_token无效
	ErrorRefreshTokenError = 40002 // refresh_token失效
	ErrorForbidden         = 40003 // 禁止访问
	ErrorNotFound          = 40004 // 请求错误,未找到该资源
	ErrorNotAllowed        = 40005 // 请求方法未允许
	ErrorTimeout           = 40008 // 请求超时
)

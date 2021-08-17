package authorize

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

const (
	UserIdKey     = "token_uid"
	AccountKey = "token_account"
	QqKey     = "token_qq"
)

func Set(r *ghttp.Request, claims map[string]interface{}) interface{} {
	r.SetParam(AccountKey, claims[AccountKey])
	r.SetParam(QqKey, claims[QqKey])
	r.SetParam(UserIdKey, claims[UserIdKey])
	return claims[UserIdKey]
}

func GetUserId(r *ghttp.Request) uint64 {
	userId := r.GetParam(UserIdKey)
	return gconv.Uint64(userId)
}

func GetAccount(r *ghttp.Request) string {
	account := r.GetParam(AccountKey)
	return gconv.String(account)
}

func GetQq(r *ghttp.Request) string {
	qq := r.GetParam(QqKey)
	return gconv.String(qq)
}

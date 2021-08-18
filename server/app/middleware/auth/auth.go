// Package auth
package auth

import (
	"dnf-game-manager/app/authorize"
	"dnf-game-manager/app/service"
	"dnf-game-manager/library/jwt"
	"dnf-game-manager/library/response"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"net/http"
	"time"
)

var JWTMiddleware *jwt.AuthJwt

// Initialization function,
// rewrite this function to customized your own JWT settings.
func init() {
	secret := g.Cfg().GetString("app.tokenSecret")
	tokenExpire := g.Cfg().GetInt64("app.tokenExpire")
	expire := time.Duration(tokenExpire)
	refreshTokenExpire := time.Duration(g.Cfg().GetInt("app.refreshTokenExpire"))
	userIdKey := authorize.UserIdKey
	authJwt, err := jwt.New(&jwt.AuthJwt{
		Key:                []byte(secret),
		Timeout:            time.Second * expire,
		RefreshTimeout:     time.Second * refreshTokenExpire,
		TokenLookup:        "header: Authorization",
		RefreshTokenLookup: "param: refreshToken",
		TokenHeadName:      "Bearer",
		TimeFunc:           time.Now,
		IdentityKey:        userIdKey,
		Authenticator:      service.Login.Login,
		LoginResponse:      LoginResponse,
		RefreshResponse:    RefreshResponse,
		Unauthorized:       Unauthorized,
		RefreshTokenError:  RefreshTokenError,
		IdentityHandler:    IdentityHandler,
		PayloadFunc:        PayloadFunc,
	})
	if err != nil {
		glog.Fatal("JWT Error:" + err.Error())
	}
	JWTMiddleware = authJwt
}

// PayloadFunc is a callback function that will be called during login.
// Using this function it is possible to add additional payload data to the webtoken.
// The data is then made available during requests via c.Get("JWT_PAYLOAD").
// Note that the payload is not encrypted.
// The attributes mentioned on jwt.io can't be used as keys for the map.
// Optional, by default no additional data will be set.
func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

// IdentityHandler sets the identity for JWT.
func IdentityHandler(r *ghttp.Request) interface{} {
	claims := jwt.ExtractClaims(r)
	return authorize.Set(r, claims)
}

// Unauthorized is used to define customized Unauthorized callback function.
func Unauthorized(r *ghttp.Request, code int, message string) {
	r.Response.WriteHeader(code)
	_ = r.Response.WriteJson(g.Map{
		"code":    code,
		"error":   response.ErrorAccessTokenError,
		"message": message,
	})
	r.ExitAll()
}
func RefreshTokenError(r *ghttp.Request, code int, message string) {
	r.Response.WriteHeader(code)
	_ = r.Response.WriteJson(g.Map{
		"code":    code,
		"error":   response.ErrorRefreshTokenError,
		"message": message,
	})
	r.ExitAll()
}

// LoginResponse is used to define customized login-successful callback function.
func LoginResponse(r *ghttp.Request, code int, token string, expire time.Time, claims g.Map) {
	//生成刷新TOKEN
	refreshToken, refreshExpire, _ := JWTMiddleware.RefreshTokenGenerator(claims)
	userIdKey := authorize.UserIdKey
	accountKey := authorize.AccountKey
	qqKey := authorize.QqKey
	_ = r.Response.WriteJson(g.Map{
		"code":                      http.StatusOK,
		"token":                     token,
		"refresh_token_expire":      refreshExpire.Format("2006-01-02 15:04:05"),
		"refresh_token_expire_unix": refreshExpire.Unix(),
		"refresh_token":             refreshToken,
		"error":                     0,
		"message":                   "登录成功",
		"expire":                    expire.Format("2006-01-02 15:04:05"),
		"expire_unix":               expire.Unix(),
		"avatar_url":                claims["avatar_url"],
		userIdKey:                   claims[userIdKey],
		accountKey:                  claims[accountKey],
		"name":                      claims[accountKey],
		qqKey:                       claims[qqKey],
	})
	r.ExitAll()
}

// RefreshResponse is used to get a new token no matter current token is expired or not.
func RefreshResponse(r *ghttp.Request, code int, token string, expire time.Time) {
	_ = r.Response.WriteJson(g.Map{
		"code":        http.StatusOK,
		"token":       token,
		"error":       0,
		"message":     "TOKEN刷新成功",
		"expire":      expire.Format("2006-01-02 15:04:05"),
		"expire_unix": expire.Unix(),
	})
	r.ExitAll()
}

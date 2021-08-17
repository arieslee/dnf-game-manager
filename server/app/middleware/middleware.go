// Package middleware
package middleware

import (
	"dnf-game-manager/app/middleware/auth"
	"github.com/gogf/gf/net/ghttp"
)

// Auth authHook is the HOOK function implements JWT logistics.
func Auth(r *ghttp.Request) {
	auth.JWTMiddleware.MiddlewareFunc()(r)
	r.Middleware.Next()
}
func CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

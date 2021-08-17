package router

import (
	"dnf-game-manager/app/middleware"
	"dnf-game-manager/app/middleware/auth"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

func Build() {
	s := g.Server("app")
	_ = s.SetConfigWithMap(g.Map{
		"serverRoot":       g.Cfg().GetString("server.root"),
		"logPath":          g.Cfg().GetString("server.logPath"),
		"dumpRouterMap":    g.Cfg().GetBool("server.dumpRouterMap"),
		"errorLogPattern":  "error-{Ymd}.log",
		"accessLogEnabled": g.Cfg().GetBool("server.accessLogEnabled"),
		"accessLogPattern": "access-{Ymd}.log",
	})
	s.Use(middleware.CORS)
	s.BindHandler("POST:/v1/login", auth.JWTMiddleware.LoginHandler)
	port := g.Cfg().GetInt("server.port")
	s.SetPort(port)
	fmt.Println("backend http 服务启动...127.0.0.1:" + gconv.String(port))
	s.Run()
}
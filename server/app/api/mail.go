package api

import (
	"dnf-game-manager/app/service"
	"dnf-game-manager/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var Mail = &mailApi{}

type mailApi struct {
}

func (a *mailApi) SendHandler(r *ghttp.Request) {
	err := service.Mail.Send(r)
	if err != nil {
		response.Error(r, err.Error())
	}
	response.Success(r, "")
}

package api

import (
	"dnf-game-manager/app/service"
	"dnf-game-manager/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var Item = &itemApi{}

type itemApi struct {
}

func (a *itemApi) ListHandler(r *ghttp.Request) {
	list, err := service.Item.List(r)
	if err != nil {
		response.Error(r, err.Error())
	}
	response.Success(r, "", list)
}

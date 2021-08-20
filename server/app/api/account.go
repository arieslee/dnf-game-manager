package api

import (
	"dnf-game-manager/app/service"
	"dnf-game-manager/library/response"
	"github.com/gogf/gf/net/ghttp"
)

var Account = &accountApi{}

type accountApi struct {
}

func (a *accountApi) ListHandler(r *ghttp.Request) {
	list, err := service.Account.List(r)
	if err != nil {
		response.Error(r, err.Error())
	}
	response.Success(r, "", list)
}
func (a *accountApi) AccountDeleteHandler(r *ghttp.Request) {
	err := service.Account.DeleteAccount(r)
	if err != nil {
		response.Error(r, err.Error())
	}
	response.Success(r, "删除成功")
}
func (a *accountApi) RoleListHandler(r *ghttp.Request) {
	list, err := service.Account.RoleList(r)
	if err != nil {
		response.Error(r, err.Error())
	}
	response.Success(r, "", list)
}

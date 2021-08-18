package service

import (
	"dnf-game-manager/app/authorize"
	"dnf-game-manager/app/dao"
	"dnf-game-manager/app/request"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Login = &loginService{

}

type loginService struct {

}

func (s *loginService)Login(r *ghttp.Request) (interface{},error) {
	req := &request.LoginInput{}
	if err := r.Parse(req);err != nil{
		return nil,err
	}
	row,_:=dao.Accounts.Fields(dao.Accounts.Columns).Where("accountname=?", req.Account).One()
	if row == nil{
		return nil,gerror.NewCode(40004, "登录账号不存在")
	}
	pwdHash,_:=gmd5.EncryptString(req.Password)
	if row.Password != pwdHash{
		return nil,gerror.NewCode(40005, "登录密码错误")
	}
	if row.Qq!="GM_master"{
		return nil,gerror.NewCode(40003, "无权登录")
	}
	userIdKey := authorize.UserIdKey
	accountKey := authorize.AccountKey
	qqKey := authorize.QqKey
	result := g.Map{
		accountKey:req.Account,
		userIdKey:row.UID,
		qqKey:row.Qq,
	}
	return result, nil
}
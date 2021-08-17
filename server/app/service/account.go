package service

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

var Account = &accountService{}

type accountService struct {

}

func (s *accountService) List(r *ghttp.Request) {

}

func (s *accountService) RoleList(r *ghttp.Request) (gdb.Result, error) {
	uid := r.GetUint64("uid")
	where:="m_id=? AND delete_flag!=1"
	params := g.Slice{uid}
	list,err := g.DB().Schema("taiwan_cain").
		Model("charac_info").
		Fields("charac_no,charac_name,job,lev,grow_type").
		Where(where,params).
		Order("charac_no ASC").
		All()
	if err != nil{
		return nil,err
	}
	return list, nil
}